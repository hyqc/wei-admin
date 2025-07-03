package global

import (
	"admin/pkg/utils"
)

func Init() error {
	for _, call := range initConfigCall {
		if err := call(); err != nil {
			utils.PrintfLn("init config error: %v", err.Error())
			return err
		}
	}
	return nil
}
