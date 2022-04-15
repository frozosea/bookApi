package review_test

import (
	"books/internal/books"
	"books/internal/review"
	"books/pkg/models"
	"books/settings"
	"books/tests/integrations/utils"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func setupReviewRepository(db *sql.DB) *review.Repository {
	repo := review.Repository{Db: db}
	return &repo
}
func getRandomBook(db *sql.DB) (*models.Book, error) {
	repo := books.Repository{Db: db}
	return repo.GetRandomBook()
}

func WriteReviewTest(t *testing.T, repo *review.Repository, userId, bookId int, review string) int {
	reviewId, err := repo.WriteReview(userId, bookId, review)
	assert.Empty(t, err)
	assert.NotEqual(t, reviewId, 0)
	return reviewId
}
func DeleteReviewTest(repo *review.Repository, reviewId int) bool {
	return repo.DeleteReview(reviewId)
}
func UpdateReviewTest(repo *review.Repository, reviewId int, newReview string) bool {
	return repo.UpdateReview(reviewId, newReview)
}
func GetReviewTest(t *testing.T, repo *review.Repository, bookId int) {
	reviews, err := repo.GetReviews(bookId)
	assert.Empty(t, err)
	for _, review := range reviews {
		assert.NotEmpty(t, review.Id)
		assert.NotEmpty(t, review.Review)
		assert.NotEmpty(t, review.ReviewId)
		assert.NotEmpty(t, review.Username)
		assert.NotEmpty(t, review.FirstName)
		assert.NotEmpty(t, review.LastName)
	}
}
func Test(t *testing.T) {
	utils.LoadEnv()
	const (
		userId    = 1
		review    = "this is test review for some books"
		newReview = "this is new test review for some books(update review)"
	)
	db, err := settings.GetDatabase()
	randomBook, exc := getRandomBook(db)
	bookId := randomBook.BookId
	repo := setupReviewRepository(db)
	if exc != nil {
		log.Fatal(exc)
	}
	if err != nil {
		log.Fatal(err)
		return
	}
	t.Run("write ,delete, update review test", func(t *testing.T) {
		reviewId := WriteReviewTest(t, repo, userId, bookId, review)
		updateReviewOk := UpdateReviewTest(repo, reviewId, newReview)
		assert.Equal(t, updateReviewOk, true)
		ok := DeleteReviewTest(repo, reviewId)
		assert.Equal(t, ok, true)
	})
	t.Run("get reviews by book id test", func(t *testing.T) {
		GetReviewTest(t, repo, bookId)
	})
}
