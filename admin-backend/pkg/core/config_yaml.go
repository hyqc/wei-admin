package core

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Yaml struct {
}

func (Yaml) Read(path string) ([]byte, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (Yaml) Decode(data []byte, conf interface{}) interface{} {
	err := yaml.Unmarshal(data, conf)
	if err != nil {
		return err
	}
	return nil
}
