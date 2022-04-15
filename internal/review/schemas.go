package review

type WriteReviewSchema struct {
	BookID int    `json:"book_id" form:"bookID"`
	Review string `json:"review" form:"Review"`
}
type UpdateReviewSchema struct {
	ReviewId  int    `json:"review_id" form:"ReviewId"`
	NewReview string `json:"new_review" form:"NewReview"`
}
type BookId struct {
	BookId int `json:"book_id" form:"BookID"`
}
type DeleteReviewSchema struct {
	ReviewId int `json:"review_id" form:"ReviewId"`
}
type WriteReviewResponseSchema struct {
	ReviewId int `json:"review_id"`
}
