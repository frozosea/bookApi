package rating

type SetRatingSchema struct {
	BookId int `json:"book_id" binding:"required" form:"BookId"`
	Rating int `json:"rating" binding:"required" form:"Rating"`
}
