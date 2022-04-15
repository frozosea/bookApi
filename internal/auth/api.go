package auth

import (
	"github.com/gin-gonic/gin"
)

type Api struct {
	Service *Service
}

func (s *Api) Login(ctx *gin.Context) {
	var schema UserAuth
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		tokens, exc := s.Service.Login(schema.Username, schema.Password)
		if exc != nil {
			ctx.JSON(401, gin.H{`err`: exc.Error()})
			return
		}
		ctx.Header(`Authorization`, tokens.AccessToken)
		ctx.JSON(200, tokens)
	} else {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
func (s *Api) Register(ctx *gin.Context) {
	var schema UserRegister
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		username := schema.Username
		password := schema.Password
		firstName := schema.FirstName
		lastName := schema.Lastname
		tokens, exc := s.Service.RegisterUser(firstName, lastName, username, password)
		if exc != nil {
			ctx.JSON(500, gin.H{`err`: exc.Error()})
			return
		}
		ctx.JSON(201, tokens)
	} else {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
