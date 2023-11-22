package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitConfig(t *testing.T) {
	conf := &Config{}
	err := conf.Init()
	assert.Nil(t, err, err)
	fmt.Println(conf.Server)
}
