package books

import (
	"books/internal/books"
	"books/settings"
	"books/tests/integrations/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func testCreateBook(t *testing.T, repo *books.Repository, BookAuthor, BookTitle string, YearOfRelease int, CoverUrl, Description string) int {
	newBook, exc := repo.CreateBook(BookAuthor, BookTitle, YearOfRelease, CoverUrl, Description)
	assert.Empty(t, exc)
	assert.Equal(t, newBook.BookAuthor, BookAuthor, `assertion error book author const != returning book's author'`)
	assert.Equal(t, newBook.BookTitle, BookTitle, `assertion error book title const != returning book's title'`)
	assert.Equal(t, newBook.YearOfRelease, YearOfRelease, `assertion error book year of release const != returning book's year of release'`)
	assert.Equal(t, newBook.CoverUrl, CoverUrl, `assertion error book cover url const != returning book's url'`)
	assert.Equal(t, newBook.Description, Description, `assertion error book description const != returning book's description'`)
	return newBook.BookId
}
func GetBookByIdTest(t *testing.T, repo books.Repository, bookId int, BookAuthor, BookTitle string, YearOfRelease int, CoverUrl, Description string) {
	resultBook, err := repo.GetBookById(bookId)
	assert.Empty(t, err)
	assert.Equal(t, resultBook.BookAuthor, BookAuthor)
	assert.Equal(t, resultBook.BookTitle, BookTitle)
	assert.Equal(t, resultBook.YearOfRelease, YearOfRelease)
	assert.Equal(t, resultBook.CoverUrl, CoverUrl)
	assert.Equal(t, resultBook.Description, Description)
}

func GetBookByAuthorTest(t *testing.T, repo books.Repository, bookId int, BookAuthor, BookTitle string, YearOfRelease int, CoverUrl, Description string) {
	Books, err := repo.GetAllBooksByAuthor(BookAuthor)
	assert.Empty(t, err)
	for _, resultBook := range Books {
		if resultBook.BookId == bookId {
			assert.Equal(t, resultBook.BookAuthor, BookAuthor)
			assert.Equal(t, resultBook.BookTitle, BookTitle)
			assert.Equal(t, resultBook.YearOfRelease, YearOfRelease)
			assert.Equal(t, resultBook.CoverUrl, CoverUrl)
			assert.Equal(t, resultBook.Description, Description)
		}
	}
}
func UpdateAuthorTest(t *testing.T, bookId int, repo *books.Repository) {
	const NewAuthor = "newAuthorTest"
	result, err := repo.UpdateAuthor(bookId, NewAuthor)
	assert.Empty(t, err)
	assert.Equal(t, result, true)
}
func UpdateYearOfReleaseTest(t *testing.T, bookId int, repo *books.Repository) {
	const NewYearOfRelease = 2022
	result, err := repo.UpdateYearOfRelease(bookId, NewYearOfRelease)
	assert.Empty(t, err)
	assert.Equal(t, result, true)
}

func UpdateTitleTest(t *testing.T, bookId int, repo *books.Repository) {
	const newBookTitle = "newBookTitleTest"
	result, err := repo.UpdateBookTitle(bookId, newBookTitle)
	assert.Empty(t, err)
	assert.Equal(t, result, true)

}
func UpdateCoverUrlTest(t *testing.T, bookId int, repo *books.Repository) {
	const newCoverUrl = "newCoverUrl.com"
	result, err := repo.UpdateCover(bookId, newCoverUrl)
	assert.Empty(t, err)
	assert.Equal(t, result, true)
}
func UpdateDescriptionTest(t *testing.T, bookId int, repo *books.Repository) {
	const newDescription = "this is updated description "
	result, err := repo.UpdateDescription(bookId, newDescription)
	assert.Empty(t, err)
	assert.Equal(t, result, true)
}
func GetRandomBookTest(t *testing.T, repo *books.Repository) {
	book, err := repo.GetRandomBook()
	assert.Empty(t, err)
	assert.NotEmpty(t, book.BookId)
	assert.NotEmpty(t, book.BookAuthor)
	assert.NotEmpty(t, book.BookTitle)
	assert.NotEmpty(t, book.YearOfRelease)
	assert.NotEmpty(t, book.Description)
	assert.NotEmpty(t, book.CoverUrl)
}
func GetAllBooksTest(t *testing.T, repo *books.Repository) {
	allBooks, err := repo.GetAllBooks()
	assert.Empty(t, err)
	for _, book := range allBooks {
		assert.NotEmpty(t, book.BookAuthor)
		assert.NotEmpty(t, book.BookId)
		assert.NotEmpty(t, book.BookTitle)
		assert.NotEmpty(t, book.YearOfRelease)
		assert.NotEmpty(t, book.CoverUrl)
		assert.NotEmpty(t, book.Description)
	}
}
func Test(t *testing.T) {
	utils.LoadEnv()
	db, err := settings.GetDatabase()
	if err != nil {
		log.Fatalf(`can not connect to database err: %s`, err)
	}
	repo := books.Repository{Db: db}
	const (
		BookAuthor    = "testBookAuthor"
		BookTitle     = "testBookTitle"
		YearOfRelease = 2022
		CoverUrl      = "https://example.com"
		Description   = "this is test book's description"
	)
	t.Run("book test", func(t *testing.T) {
		bookId := testCreateBook(t, &repo, BookAuthor, BookTitle, YearOfRelease, CoverUrl, Description)
		GetBookByIdTest(t, repo, bookId, BookAuthor, BookTitle, YearOfRelease, CoverUrl, Description)
		GetBookByAuthorTest(t, repo, bookId, BookAuthor, BookTitle, YearOfRelease, CoverUrl, Description)
		GetAllBooksTest(t, &repo)
		GetRandomBookTest(t, &repo)
		UpdateAuthorTest(t, bookId, &repo)
		UpdateYearOfReleaseTest(t, bookId, &repo)
		UpdateTitleTest(t, bookId, &repo)
		UpdateCoverUrlTest(t, bookId, &repo)
		UpdateDescriptionTest(t, bookId, &repo)

	})
}
