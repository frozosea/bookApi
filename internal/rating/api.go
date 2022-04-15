package rating

import (
	"github.com/gin-gonic/gin"
)

type Api struct {
	Service ServiceRating
}

func (a *Api) SetRating(ctx *gin.Context) {
	var schema SetRatingSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		ok, exc := a.Service.SetRating(schema.BookId, schema.Rating)
		if !ok || exc != nil {
			ctx.JSON(500, gin.H{`err`: exc.Error()})
		}
		ctx.JSON(201, gin.H{`success`: true})
	} else {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
