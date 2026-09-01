// Package captcha 基于 github.com/mojocn/base64Captcha 的验证码封装。
//
// 设计要点：
//  1. 按「场景」(Scene) 管理，新增业务验证码（注册、找回密码等）只需定义 Scene 并加配置；
//  2. 存储 (Store) 可替换：单机部署用内存存储，集群部署用 Redis 存储，由外部注入；
//  3. 每个场景独立配置位数/尺寸/有效期，并独立持有存储实例（便于不同场景不同有效期）。
package captcha

import (
	"fmt"
	"strings"
	"time"

	"github.com/mojocn/base64Captcha"
)

// Scene 验证码使用场景
type Scene string

const (
	// SceneLogin 登录验证码
	SceneLogin Scene = "login"
)

// 默认值（配置未设置时使用）
const (
	DefaultLength = 4
	DefaultWidth  = 120
	DefaultHeight = 44
	DefaultExpire = 5 * time.Minute
)

// Chars 验证码字符池（去掉易混淆的 0 / 1 / I / O）
const Chars = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// 干扰强度：优先保证肉眼可辨识——不叠加干扰线，仅保留极少量噪点
const (
	// noiseCount 噪点数量
	noiseCount = 1
	// showLineOptions 干扰线选项：0 表示不绘制干扰线
	// （base64Captcha.OptionShowSlimeLine / OptionShowSineLine 会显著降低可读性）
	showLineOptions = 0
)

// imagePrefix base64Captcha 返回纯 base64 图片内容，前端展示需补充 data URL 前缀
const imagePrefix = "data:image/png;base64,"

// Store 验证码存储，与 base64Captcha.Store 接口保持一致
type Store = base64Captcha.Store

// Config 单个场景的验证码配置
type Config struct {
	Length int           // 验证码位数
	Expire time.Duration // 有效期
	Width  int           // 图片宽度
	Height int           // 图片高度
}

func (c Config) fix() Config {
	if c.Length <= 0 {
		c.Length = DefaultLength
	}
	if c.Expire <= 0 {
		c.Expire = DefaultExpire
	}
	if c.Width <= 0 {
		c.Width = DefaultWidth
	}
	if c.Height <= 0 {
		c.Height = DefaultHeight
	}
	return c
}

type sceneEntry struct {
	config Config
	store  Store
}

// Manager 验证码管理器
type Manager struct {
	scenes map[Scene]*sceneEntry
}

// New 创建验证码管理器。
// storeFactory 用于为每个场景创建存储实例（单机返回内存存储、集群返回 Redis 存储），
// 创建失败（如配置了 Redis 存储但 Redis 不可用）时直接返回错误，避免静默降级；
// scenes 为各场景配置，为空时自动注册默认登录场景。
func New(storeFactory func(Scene, Config) (Store, error), scenes map[Scene]Config) (*Manager, error) {
	if len(scenes) == 0 {
		scenes = map[Scene]Config{SceneLogin: {}}
	}
	m := &Manager{scenes: make(map[Scene]*sceneEntry, len(scenes))}
	for scene, cfg := range scenes {
		cfg = cfg.fix()
		store, err := storeFactory(scene, cfg)
		if err != nil {
			return nil, fmt.Errorf("captcha: 场景 %q 创建存储失败: %w", scene, err)
		}
		m.scenes[scene] = &sceneEntry{
			config: cfg,
			store:  store,
		}
	}
	return m, nil
}

// Scenes 返回已注册的场景列表
func (m *Manager) Scenes() []Scene {
	list := make([]Scene, 0, len(m.scenes))
	for scene := range m.scenes {
		list = append(list, scene)
	}
	return list
}

// Generate 生成指定场景的验证码，返回验证码 ID 与可直接用于 img src 的 data URL
func (m *Manager) Generate(scene Scene) (id string, dataURL string, err error) {
	id, dataURL, _, err = m.GenerateWithAnswer(scene)
	return id, dataURL, err
}

// GenerateWithAnswer 同 Generate，额外返回验证码明文（仅用于调试/测试环境记录日志，禁止响应给前端）
func (m *Manager) GenerateWithAnswer(scene Scene) (id string, dataURL string, answer string, err error) {
	entry, ok := m.scenes[scene]
	if !ok {
		return "", "", "", fmt.Errorf("captcha: 未配置的验证码场景 %q", scene)
	}
	driver := base64Captcha.NewDriverString(
		entry.config.Height,
		entry.config.Width,
		noiseCount, // 噪点数量：少量，不影响辨识
		showLineOptions,
		entry.config.Length,
		Chars,
		nil, // 背景色：使用库默认值
		nil, // 字体存储：为 nil 时使用库内置字体
		nil, // 字体名：为空时使用内置默认字体
	)
	// 返回值依次为 验证码ID、base64 图片、明文答案、错误
	id, b64s, answer, err := base64Captcha.NewCaptcha(driver, entry.store).Generate()
	if err != nil {
		return "", "", "", err
	}
	// 库返回的图片内容已包含 data URL 前缀时不再重复拼接
	if strings.HasPrefix(b64s, "data:image") {
		return id, b64s, answer, nil
	}
	return id, imagePrefix + b64s, answer, nil
}

// Verify 校验指定场景的验证码：忽略大小写，校验通过后立即失效（一次性）
func (m *Manager) Verify(scene Scene, id, code string) bool {
	entry, ok := m.scenes[scene]
	if !ok {
		return false
	}
	code = strings.ToUpper(strings.TrimSpace(code))
	if id == "" || code == "" {
		return false
	}
	return entry.store.Verify(id, code, true)
}
