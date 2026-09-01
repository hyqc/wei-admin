package captcha

import (
	"errors"
	"strings"
	"testing"
	"time"
)

// recordStore 测试用存储：记录写入内容，便于断言生成结果
type recordStore struct {
	values map[string]string
}

func (s *recordStore) Set(id, value string) error {
	if s.values == nil {
		s.values = map[string]string{}
	}
	s.values[id] = value
	return nil
}

func (s *recordStore) Get(id string, clear bool) string {
	v := s.values[id]
	if clear {
		delete(s.values, id)
	}
	return v
}

func (s *recordStore) Verify(id, answer string, clear bool) bool {
	return s.Get(id, clear) == answer
}

func newTestManager(scenes map[Scene]Config) (*Manager, *recordStore) {
	rs := &recordStore{values: map[string]string{}}
	m, err := New(func(_ Scene, _ Config) (Store, error) { return rs, nil }, scenes)
	if err != nil {
		panic(err)
	}
	return m, rs
}

func TestStoreFactoryError(t *testing.T) {
	// 存储创建失败时不应静默降级
	if _, err := New(func(_ Scene, _ Config) (Store, error) {
		return nil, errors.New("redis unavailable")
	}, nil); err == nil {
		t.Fatal("store factory error should be returned")
	}
}

func TestGenerateAndVerify(t *testing.T) {
	m, rs := newTestManager(map[Scene]Config{SceneLogin: {Length: 6, Expire: time.Minute, Width: 120, Height: 44}})
	id, dataURL, err := m.Generate(SceneLogin)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}
	if id == "" {
		t.Fatal("captcha id is empty")
	}
	if !strings.HasPrefix(dataURL, "data:image/png;base64,") {
		t.Fatalf("unexpected data url prefix: %s", dataURL[:32])
	}
	code, ok := rs.values[id]
	if !ok {
		t.Fatal("captcha not stored")
	}
	if len(code) != 6 {
		t.Fatalf("code length = %d, want 6", len(code))
	}
	for _, c := range code {
		if !strings.ContainsRune(Chars, c) {
			t.Fatalf("char %q not in charset", c)
		}
	}
	// 忽略大小写
	if !m.Verify(SceneLogin, id, strings.ToLower(code)) {
		t.Fatal("verify should ignore case")
	}
	// 一次性：再次校验应失败
	if m.Verify(SceneLogin, id, code) {
		t.Fatal("captcha should be one-time")
	}
}

func TestDefaultSceneRegistered(t *testing.T) {
	// 未配置任何场景时，应自动注册默认登录场景
	m, rs := newTestManager(nil)
	if len(m.Scenes()) != 1 || m.Scenes()[0] != SceneLogin {
		t.Fatalf("scenes = %v, want [login]", m.Scenes())
	}
	id, _, err := m.Generate(SceneLogin)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}
	if len(rs.values[id]) != DefaultLength {
		t.Fatalf("code length = %d, want %d", len(rs.values[id]), DefaultLength)
	}
}

func TestUnknownScene(t *testing.T) {
	m, _ := newTestManager(nil)
	if _, _, err := m.Generate(Scene("register")); err == nil {
		t.Fatal("unknown scene should return error")
	}
	if m.Verify(Scene("register"), "id", "code") {
		t.Fatal("unknown scene verify should fail")
	}
}

func TestVerifyEmptyInput(t *testing.T) {
	m, rs := newTestManager(nil)
	id, _, err := m.Generate(SceneLogin)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}
	if m.Verify(SceneLogin, "", rs.values[id]) || m.Verify(SceneLogin, id, "  ") {
		t.Fatal("empty id or code should fail")
	}
}

func TestMemoryStoreExpire(t *testing.T) {
	store := NewMemoryStore(50 * time.Millisecond)
	store.Set("expire-id", "ABCD")
	if got := store.Get("expire-id", false); got != "ABCD" {
		t.Fatalf("Get = %q, want ABCD", got)
	}
	time.Sleep(80 * time.Millisecond)
	if got := store.Get("expire-id", false); got != "" {
		t.Fatalf("expired value should be empty, got %q", got)
	}
}
