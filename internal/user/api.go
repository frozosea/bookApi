package user

import (
	"books/internal/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

type Api struct {
	Service   ServiceInterface
	SecretKey string
}

func (a *Api) UpdateFirstName(c *gin.Context) {
	var schema UpdateFirstNameSchema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(400, gin.H{`err`: err.Error()})
		return
	}
	rawToken := c.GetHeader(`Authorization`)
	token := strings.Split(rawToken, " ")
	userID, exc := auth.GetUserIdByJwtToken(token[1], a.SecretKey)
	if exc != nil {
		c.JSON(401, gin.H{`err`: exc.Error()})
		return
	}
	ok, exception := a.Service.UpdateFirstName(userID, schema.NewFirstName)
	if !ok || exception != nil {
		c.JSON(500, gin.H{`success`: false, `err`: exception.Error()})
		return
	}
	c.JSON(200, gin.H{`success`: true})
}
func (a *Api) UpdateLastname(c *gin.Context) {
	var schema UpdateLastNameSchema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(400, gin.H{`err`: err.Error()})
		return
	}

	rawToken := c.GetHeader(`Authorization`)
	token := strings.Split(rawToken, " ")
	userID, exc := auth.GetUserIdByJwtToken(token[1], a.SecretKey)
	if exc != nil {
		c.JSON(401, gin.H{`err`: exc.Error()})
		return
	}
	ok, exception := a.Service.UpdateLastName(userID, schema.NewLastName)
	if !ok || exception != nil {
		c.JSON(500, gin.H{`success`: false, `err`: exception.Error()})
		return
	}
	c.JSON(200, gin.H{`success`: true})
}
func (a *Api) UpdateUsername(c *gin.Context) {
	var schema UpdateUsernameSchema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(400, gin.H{`err`: err.Error()})
		return
	}
	rawToken := c.GetHeader(`Authorization`)
	token := strings.Split(rawToken, " ")
	userID, exc := auth.GetUserIdByJwtToken(token[1], a.SecretKey)
	if exc != nil {
		c.JSON(401, gin.H{`err`: exc.Error()})
		return
	}
	ok, exception := a.Service.UpdateUsername(userID, schema.NewUsername)
	if !ok || exception != nil {
		c.JSON(500, gin.H{`success`: false, `err`: exception.Error()})
		return
	}
	c.JSON(200, gin.H{`success`: true})
}

//UpdatePassword godoc
//@Summary Update password
//@Produce json
//@Param id path integer true "User ID"
func (a *Api) UpdatePassword(c *gin.Context) {
	var schema UpdatePasswordSchema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(400, gin.H{`err`: err.Error()})
		return
	}
	rawToken := c.GetHeader(`Authorization`)
	token := strings.Split(rawToken, " ")
	userID, exc := auth.GetUserIdByJwtToken(token[1], a.SecretKey)
	if exc != nil {
		c.JSON(401, gin.H{`err`: exc.Error()})
		return
	}
	ok, exception := a.Service.UpdatePassword(userID, schema.NewPassword)
	if !ok || exception != nil {
		c.JSON(500, gin.H{`success`: false, `err`: exception.Error()})
	}
	c.JSON(200, gin.H{`success`: true})
}
func (a *Api) GetInfoAboutUser(c *gin.Context) {
	var schema *GetInfoAboutUserSchema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(400, gin.H{`err`: err.Error()})
		return
	}
	user, exc := a.Service.GetInfoAboutUser(schema.UserId)
	if exc != nil {
		c.JSON(500, gin.H{`success`: false, `err`: exc.Error()})
		return
	}
	c.JSON(200, user)
}
