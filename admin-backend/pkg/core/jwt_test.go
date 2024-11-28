package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenCreate(t *testing.T) {
	option := CustomClaimsOption{
		AccountId:     1,
		ExpireSeconds: 604800,
		UUID:          1,
		Secret:        "MHcCAQEEIEiE+ONCLujTc4ibxd9qtfVRGlrSX2A0qXYVzrw4hUREoAoGCCqGSM49AwEHoUQDQgAEisyae3uURSutgrUTw8r14dC1Wtjflq9l3D75jk+9JQprkXQB7TFBwRCdco3/h+oAQrozhqvgKG+nulGIoOSpaQ==",
	}
	token, err := JWTCreate(option)
	assert.Nil(t, err, "错误")
	fmt.Println("token: ", token)

	res, err := JWTCheck(token, option.Secret)
	assert.Nil(t, err, "错误")
	fmt.Println(res.ID, res.AdminID, res.Subject, res, res.ExpiresAt)

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzMzODM2MzcsImlhdCI6MTczMjc3ODgzNywianRpIjoiNTQyMzIyMjI3Mjk5NzcwMzY5IiwiYWRtaW5faWQiOjF9.nx7wo5J2api_Mil99cyYCTDOTbn2n4GCwDU0BdDHGME"
	res, err = JWTCheck(token, option.Secret)
	assert.Nil(t, err, "错误")
	fmt.Println(res.ID, res.AdminID, res.Subject, res, res.ExpiresAt)
}
