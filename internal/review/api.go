package review

import (
	"books/internal/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

type Api struct {
	Service   ServiceInterface
	SecretKey string
}

func (s *Api) getUserIDFromContext(ctx *gin.Context) int {
	rawToken := ctx.GetHeader(`Authorization`)
	token := strings.Split(rawToken, " ")
	id, exc := auth.GetUserIdByJwtToken(token[1], s.SecretKey)
	if exc != nil {
		ctx.JSON(401, gin.H{`err`: exc.Error()})
		return 1
	}
	return id
}
func (s *Api) WriteNewReview(ctx *gin.Context) {
	var schema WriteReviewSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		userID := s.getUserIDFromContext(ctx)
		reviewId, exc := s.Service.WriteReview(userID, schema.BookID, schema.Review)
		if exc != nil {
			ctx.JSON(500, gin.H{`err`: exc.Error()})
		}
		ctx.JSON(201, gin.H{`review_id`: reviewId})
	} else {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
func (s *Api) UpdateWriteReview(ctx *gin.Context) {
	var schema UpdateReviewSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		_ = s.getUserIDFromContext(ctx)
		go s.Service.UpdateReview(schema.ReviewId, schema.NewReview)
		ctx.JSON(200, gin.H{`success`: true})
	} else {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
func (s *Api) GetAllReviewsByBookID(ctx *gin.Context) {
	var schema BookId
	if err := ctx.ShouldBindJSON(&schema); err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	result, err := s.Service.GetReviews(schema.BookId)
	if err != nil {
		ctx.JSON(500, gin.H{`err`: err.Error()})
	}
	ctx.JSON(200, result)
}
func (s *Api) DeleteReview(ctx *gin.Context) {
	var schema DeleteReviewSchema
	if err := ctx.ShouldBindJSON(&schema); err != nil {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
	_ = s.getUserIDFromContext(ctx)
	go s.Service.DeleteReview(schema.ReviewId)
	ctx.JSON(200, gin.H{`success`: true})
}
