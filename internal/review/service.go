package review

import "books/pkg/models"

type ServiceInterface interface {
	WriteReview(userId, bookId int, review string) (int, error)
	UpdateReview(reviewId int, newReview string) bool
	GetReviews(bookID int) ([]models.UserReview, error)
	DeleteReview(reviewId int) bool
}
type Service struct {
	Repository Repository
}

//WriteReview ...
func (s *Service) WriteReview(userId, bookId int, review string) (int, error) {
	return s.Repository.WriteReview(userId, bookId, review)
}

//UpdateReview ...
func (s *Service) UpdateReview(reviewId int, newReview string) bool {
	return s.Repository.UpdateReview(reviewId, newReview)
}

//GetReviews get all reviews by book id
func (s *Service) GetReviews(bookID int) ([]models.UserReview, error) {
	return s.Repository.GetReviews(bookID)
}

//DeleteReview delete Review by id
func (s *Service) DeleteReview(reviewId int) bool {
	return s.Repository.DeleteReview(reviewId)
}
