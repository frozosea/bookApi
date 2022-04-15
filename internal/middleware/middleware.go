package middleware

import (
	"books/internal/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthMiddleware struct {
	SecretKey string
}

func (b *AuthMiddleware) Auth(ctx *gin.Context) {
	authHeader := ctx.GetHeader(`Authorization`)
	if authHeader == "" {
		ctx.AbortWithStatus(401)
		return
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 {
		ctx.AbortWithStatus(401)
		return
	}
	_, exc := auth.GetUserIdByJwtToken(authParts[1], b.SecretKey)
	if exc != nil {
		ctx.AbortWithStatus(401)
	}
}
