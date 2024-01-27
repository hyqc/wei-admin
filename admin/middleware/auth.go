package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("authorization")
		token, err := jwt.ParseWithClaims()
		jwt.ClaimsValidator()
	}
}
