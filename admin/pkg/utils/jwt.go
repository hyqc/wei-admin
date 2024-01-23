package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	AdminID   int    `json:"admin_id"`
	AdminName string `json:"admin_name"`
}

func TokenCreate(accountId int, accountName string, uuid int64) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodES256, CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    accountName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        strconv.FormatInt(uuid, 10),
		},
		AdminID:   accountId,
		AdminName: accountName,
	})
}

func TokenValidate(token string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodES256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return nil, nil
	})
}
