package config

import (
	json2 "encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestResetJsonOrigin(t *testing.T) {
	type A struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Max  int    `json:"max"`
	}

	type B struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	a := &A{
		Id:   1,
		Name: "1",
		Max:  1,
	}

	b := &B{
		Id:   2,
		Name: "2",
	}

	aBody, _ := json2.Marshal(a)
	fmt.Println(string(aBody))
	bBody, _ := json2.Marshal(b)
	_ = json2.Unmarshal(bBody, a)
	assert.Equal(t, 1, a.Max)

	err := resetJsonOrigin(bBody, a)
	fmt.Println(a)
	assert.Nil(t, err, "错误")
	assert.Equal(t, 0, a.Max)
}

type testCf struct {
	Server Server      `json:"server"`
	Nacos  NacosConfig `json:"nacos"`
	Broker Broker      `json:"broker"`
	JWT    Jwt         `json:"jwt"`
}

func (n *testCf) Original() interface{} {
	return n
}

func (n *testCf) Handle() error {
	return nil
}

func TestInit(t *testing.T) {
	cf := &testCf{}
	err := Init(cf)
	fmt.Println("=========----------- ", cf)
	assert.Nil(t, err, "错误")
}

type Mail struct {
	Id        int32  `json:"id"`
	MailTitle string `json:"mailTitle"`
	MailText  string `json:"mailText"`
}

type mailConf struct {
	rw       *sync.RWMutex
	original []*Mail
	m        map[int32]*Mail
}

func (m *mailConf) Handle() error {
	m.rw.Lock()
	defer m.rw.Unlock()
	m.m = make(map[int32]*Mail)
	for _, val := range m.original {
		m.m[val.Id] = val
	}
	return nil
}

func (m *mailConf) Original() interface{} {
	m.rw.RLock()
	defer m.rw.RUnlock()
	return &m.original
}

func newMailConf(nc *NacosConfig) (*mailConf, error) {
	cf := &mailConf{
		rw:       &sync.RWMutex{},
		original: make([]*Mail, 0),
		m:        map[int32]*Mail{},
	}
	err := LoadByNacosDataId(nc, cf, "mail_config.json", "mail_config")
	return cf, err
}

func TestLoadByNacosDataId(t *testing.T) {
	cf := &testCf{}
	if err := Init(cf); err != nil {
		t.Fatal(err)
		return
	}
	cc, err := newMailConf(Nacos)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(cc.m)
}
