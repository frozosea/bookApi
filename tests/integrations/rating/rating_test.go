package rating_test

import (
	"books/internal/books"
	"books/internal/rating"
	"books/pkg/models"
	"books/settings"
	"books/tests/integrations/utils"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
)

func setupRatingRepo(db *sql.DB) *rating.Repository {
	repo := rating.Repository{Db: db}
	return &repo
}
func getRandomBook(db *sql.DB) (*models.Book, error) {
	repo := books.Repository{Db: db}
	return repo.GetRandomBook()
}
func SetRatingTest(t *testing.T, repo *rating.Repository, bookId int) bool {
	var randomRatingArray []int
	for i := 0; i < 100; i++ {
		randomRatingArray = append(randomRatingArray, rand.Intn(5))
	}
	for _, item := range randomRatingArray {
		ok, err := repo.SetRating(bookId, item)
		assert.Empty(t, err)
		assert.Equal(t, ok, true)
	}
	return true
}
func Test(t *testing.T) {
	utils.LoadEnv()
	db, err := settings.GetDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	ratingRepo := setupRatingRepo(db)
	randomBook, exc := getRandomBook(db)
	if exc != nil {
		log.Fatal(exc.Error())
	}
	bookId := randomBook.BookId
	t.Run("set rating test", func(t *testing.T) {
		SetRatingTest(t, ratingRepo, bookId)
	})

}
