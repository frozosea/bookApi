package books

import (
	"books/internal/auth"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func Middleware(ctx *gin.Context) {
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
	_, exc := auth.GetUserIdByJwtToken(authParts[1], os.Getenv(`JWT_SECRET_KEY`))
	if exc != nil {
		ctx.AbortWithStatus(401)
	}
}
