package core

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	AdminID int32 `json:"admin_id"`
}

type CustomClaimsOption struct {
	AccountId     int32  // 账号ID
	ExpireSeconds int64  // 过期秒数
	UUID          uint64 // 唯一ID
	Secret        string // 加密密钥
}

func newCustomClaims(option CustomClaimsOption) *CustomClaims {
	return &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(option.ExpireSeconds) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        strconv.FormatUint(option.UUID, 10),
		},
		AdminID: option.AccountId,
	}
}

func JWTCreate(option CustomClaimsOption) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, newCustomClaims(option)).SignedString([]byte(option.Secret))
}

func JWTCheck(t, secret string) (*CustomClaims, error) {
	result := &CustomClaims{}
	token, err := jwt.ParseWithClaims(t, result, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid == false {
		return nil, errors.New("token valid failed")
	}
	return result, err
}
