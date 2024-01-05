package uuid

import (
	"github.com/sony/sonyflake"
)

var Sonyflake *sonyflake.Sonyflake

func init() {
	initSnoyflake()
}

func initSnoyflake() {
	var st sonyflake.Settings
	Sonyflake = sonyflake.NewSonyflake(st)
}
