package config

import (
	"errors"
	"github.com/sony/sonyflake"
)

func InitSnoyflake() error {
	var st sonyflake.Settings
	AppSnoyflake = sonyflake.NewSonyflake(st)
	if AppSnoyflake == nil {
		return errors.New("snoyflake init nil")
	}
	return nil
}
