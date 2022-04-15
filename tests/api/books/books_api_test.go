package books_api_test

import (
	"books/internal/books"
	"books/pkg/models"
	api_utils "books/tests/api"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"testing"
)

const (
	BookAuthor           = "testApiBookAuthor"
	BookTitle            = "testApiBookTitle"
	YearOfRelease        = 2022
	CoverUrl             = "https://example.com"
	Description          = "this is test api book's description"
	baseApiUrl           = "http://localhost:8080"
	newBookYearOfRelease = 2023
	newDescription       = "this is new description for random book send to api request"
	newBookAuthor        = "owner of this"
	newBookTitle         = "new book title for api"
	newCoverUrl          = baseApiUrl + "newCover"
	bookId               = 1
	username             = "testApiUsername"
	password             = "testApiPassword"
)

func getHttpClient() *http.Client {
	client := http.Client{}
	return &client
}
func CreateBookTest(t *testing.T, baseUrl, accessToken string) {
	CreateBookSchema := books.CreateBookSchema{
		BookAuthor:    BookAuthor,
		BookTitle:     BookTitle,
		YearOfRelease: YearOfRelease,
		Description:   Description,
		CoverUrl:      CoverUrl}
	requestJson, err := json.Marshal(CreateBookSchema)
	var book models.Book
	assert.Empty(t, err)
	var url = fmt.Sprintf(`%s/v1/books/CreateBook`, baseUrl)
	body := SendRequestWithClient(t, "POST", url, accessToken, requestJson)
	if exc := json.Unmarshal(body, &book); exc != nil {
		log.Fatal(exc.Error())
	}
	assert.Empty(t, err)
	AssertEqualsOfBookData(t, book)
	assert.NotEmpty(t, book.BookId)

}
func AssertEqualsOfBookData(t *testing.T, book models.Book) {
	assert.Equal(t, book.BookAuthor, BookAuthor)
	assert.Equal(t, book.BookTitle, BookTitle)
	assert.Equal(t, book.YearOfRelease, YearOfRelease)
	assert.Equal(t, book.CoverUrl, CoverUrl)
	assert.Equal(t, book.Description, Description)
}
func assertBookIsNotEmpty(t *testing.T, book models.Book) {
	assert.NotEmpty(t, book.BookId)
	assert.NotEmpty(t, book.BookAuthor)
	assert.NotEmpty(t, book.BookTitle)
	assert.NotEmpty(t, book.YearOfRelease)
	assert.NotEmpty(t, book.Description)
	assert.NotEmpty(t, book.CoverUrl)
}
func GetAllBooksTest(t *testing.T, baseUrl string) {
	var AllBooksUrl = fmt.Sprintf(`%s/books/GetAllBooks`, baseUrl)
	Books := sendGetRequestToApiWithManyBooksResponse(t, AllBooksUrl, []byte(``))
	for _, book := range Books {
		assertBookIsNotEmpty(t, book)
	}
}
func sendGetRequestToApiWithManyBooksResponse(t *testing.T, url string, requestBody []byte) []models.Book {
	var Books []models.Book
	body := sendGetRequestWithJson(t, "GET", url, requestBody)
	exc := json.Unmarshal(body, &Books)
	assert.Empty(t, exc)
	return Books
}
func GetRandomBookTest(t *testing.T, baseUrl string) *models.Book {
	var randomBookUrl = fmt.Sprintf(`%s/books/random`, baseUrl)
	book := sendGetRequestToApiWithOneBookResponse(t, randomBookUrl)
	assertBookIsNotEmpty(t, book)
	return &book
}
func sendGetRequestAndGetByteBody(t *testing.T, url string) []byte {
	response, err := http.Get(url)
	assert.Empty(t, err)
	defer response.Body.Close()
	body, errs := io.ReadAll(response.Body)
	assert.Empty(t, errs)
	return body
}
func sendGetRequestWithJson(t *testing.T, requestMethod, url string, requestBody []byte) []byte {
	client := getHttpClient()
	req, err := http.NewRequest(requestMethod, url, bytes.NewBuffer(requestBody))
	assert.Empty(t, err)
	response, exc := client.Do(req)
	assert.Empty(t, exc)
	defer response.Body.Close()
	body, errs := io.ReadAll(response.Body)
	assert.Empty(t, errs)
	return body
}
func sendGetRequestToApiWithOneBookResponse(t *testing.T, url string) models.Book {
	var book models.Book
	body := sendGetRequestAndGetByteBody(t, url)
	exc := json.Unmarshal(body, &book)
	assert.Empty(t, exc)
	return book
}
func GetBookByIdTest(t *testing.T, bookId int, baseUrl string) {
	var book models.Book
	var requestSchema = books.BookIdSchema{BookId: bookId}
	requestJson, err := json.Marshal(requestSchema)
	assert.Empty(t, err)
	var BookIdUrl = fmt.Sprintf(`%s/books/GetBookByID`, baseUrl)
	body := sendGetRequestWithJson(t, "GET", BookIdUrl, requestJson)
	if exc := json.Unmarshal(body, &book); exc != nil {
		log.Fatal(exc)
	}
	assertBookIsNotEmpty(t, book)
}
func GetBooksByAuthorTest(t *testing.T, author, baseUrl string) {
	var AuthorUrl = fmt.Sprintf(`%s/books/author`, baseUrl)
	var requestSchema = books.GetAllBooksByAuthorSchema{Author: author}
	requestJson, err := json.Marshal(requestSchema)
	assert.Empty(t, err)
	Books := sendGetRequestToApiWithManyBooksResponse(t, AuthorUrl, requestJson)
	for _, book := range Books {
		assertBookIsNotEmpty(t, book)
	}
}
func SendRequestWithClient(t *testing.T, requestMethod, url string, accessToken string, requestBody []byte) []byte {
	client := getHttpClient()
	req, err := http.NewRequest(requestMethod, url, bytes.NewBuffer(requestBody))
	assert.Empty(t, err)
	req.Header.Set(`Authorization`, fmt.Sprintf(`Bearer %s`, accessToken))
	response, exc := client.Do(req)
	assert.Empty(t, exc)
	defer response.Body.Close()
	body, errs := io.ReadAll(response.Body)
	assert.Empty(t, errs)
	return body
}
func checkBaseResponse(t *testing.T, body []byte) {
	var result books.BaseResponse
	if exc := json.Unmarshal(body, &result); exc != nil {
		log.Fatal(exc.Error())
	}
	assert.Equal(t, result.Success, true)
}
func UpdateYearOfReleaseTest(t *testing.T, bookId, newYearOfRelease int, baseUrl, accessToken string) {
	updateYearOfRelease := books.UpdateYearOfReleaseSchema{BookIdSchema: books.BookIdSchema{BookId: bookId}, NewYearOfRelease: newYearOfRelease}
	var updYearUrl = fmt.Sprintf(`%s/v1/books/updateYearOfRelease`, baseUrl)
	requestBody, exc := json.Marshal(updateYearOfRelease)
	assert.Empty(t, exc)
	body := SendRequestWithClient(t, "PUT", updYearUrl, accessToken, requestBody)
	checkBaseResponse(t, body)
}
func UpdateAuthorTest(t *testing.T, bookId int, newAuthor, accessToken, baseUrl string) {
	UpdateAuthor := books.UpdateAuthorSchema{BookIdSchema: books.BookIdSchema{BookId: bookId}, NewAuthor: newAuthor}
	var updAuthorUrl = fmt.Sprintf(`%s/v1/books/updateAuthor`, baseUrl)
	requestBody, exc := json.Marshal(UpdateAuthor)
	assert.Empty(t, exc)
	body := SendRequestWithClient(t, "PUT", updAuthorUrl, accessToken, requestBody)
	checkBaseResponse(t, body)
}

func UpdateDescriptionTest(t *testing.T, bookId int, newDescription, accessToken, baseUrl string) {
	UpdateDescription := books.UpdateDescriptionSchema{BookIdSchema: books.BookIdSchema{BookId: bookId}, NewDescription: newDescription}
	var updDescriptionUrl = fmt.Sprintf(`%s/v1/books/updateDescription`, baseUrl)
	requestBody, exc := json.Marshal(UpdateDescription)
	assert.Empty(t, exc)
	body := SendRequestWithClient(t, "PUT", updDescriptionUrl, accessToken, requestBody)
	checkBaseResponse(t, body)
}
func UpdateTitleTest(t *testing.T, bookId int, baseUrl, newTitle, accessToken string) {
	newTitleSchema := books.UpdateBookTitleSchema{BookIdSchema: books.BookIdSchema{BookId: bookId}, NewBookTitle: newTitle}
	var updTitleUrl = fmt.Sprintf(`%s/v1/books/updateTitle`, baseUrl)
	requestBody, exc := json.Marshal(newTitleSchema)
	assert.Empty(t, exc)
	body := SendRequestWithClient(t, "PUT", updTitleUrl, accessToken, requestBody)
	checkBaseResponse(t, body)
}
func UpdateCoveUrlTest(t *testing.T, bookId int, baseUrl, newCoverUrl, accessToken string) {
	newCoverUrlSchema := books.UpdateCoverSchema{BookIdSchema: books.BookIdSchema{BookId: bookId}, NewCoverUrl: newCoverUrl}
	var updTitleUrl = fmt.Sprintf(`%s/v1/books/updateCover`, baseUrl)
	requestBody, exc := json.Marshal(newCoverUrlSchema)
	assert.Empty(t, exc)
	body := SendRequestWithClient(t, "PUT", updTitleUrl, accessToken, requestBody)
	checkBaseResponse(t, body)
}
func Test(t *testing.T) {
	accessToken := api_utils.Login(username, password)
	randomBook := GetRandomBookTest(t, baseApiUrl)
	t.Run("create book test", func(te *testing.T) {
		CreateBookTest(te, baseApiUrl, accessToken)
	})
	t.Run("update book's params tests", func(te *testing.T) {
		UpdateYearOfReleaseTest(te, randomBook.BookId, newBookYearOfRelease, baseApiUrl, accessToken)
		UpdateAuthorTest(te, randomBook.BookId, newBookAuthor, accessToken, baseApiUrl)
		UpdateDescriptionTest(te, randomBook.BookId, newDescription, accessToken, baseApiUrl)
		UpdateCoveUrlTest(te, randomBook.BookId, baseApiUrl, newCoverUrl, accessToken)
		UpdateTitleTest(te, randomBook.BookId, baseApiUrl, newBookTitle, accessToken)
	})
	t.Run("get books test", func(te *testing.T) {
		GetAllBooksTest(te, baseApiUrl)
		GetBooksByAuthorTest(te, BookAuthor, baseApiUrl)
		GetBookByIdTest(t, bookId, baseApiUrl)
	})
}
