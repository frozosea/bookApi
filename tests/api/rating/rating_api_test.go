package rating_api_test

import (
	"books/internal/books"
	"books/internal/rating"
	"books/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"math/rand"
	"net/http"
	"testing"
)

const baseUrl = "http://localhost:8080"

func getRandomBooks(t *testing.T) *models.Book {
	var getRandBookUrl = fmt.Sprintf(`%s/books/random`, baseUrl)
	var randomBook *models.Book
	response, exc := http.Get(getRandBookUrl)
	assert.Empty(t, exc)
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	assert.Empty(t, err)
	if exception := json.Unmarshal(body, &randomBook); exception != nil {
		log.Fatal(exception)
	}
	return randomBook
}
func sendRequestAndGetBody(t *testing.T, bookId, bookRating int) []byte {
	var requestSchema = rating.SetRatingSchema{
		Rating: bookRating, BookId: bookId,
	}
	var setRatingUrl = fmt.Sprintf(`%s/rating/setRating`, baseUrl)
	requestBody, exception := json.Marshal(requestSchema)
	client := http.Client{}
	assert.Empty(t, exception)
	req, err := http.NewRequest("POST", setRatingUrl, bytes.NewBuffer(requestBody))
	assert.Empty(t, err)
	response, exc := client.Do(req)
	assert.Empty(t, exc)
	defer response.Body.Close()
	assert.Equal(t, response.StatusCode, 201)
	body, e := io.ReadAll(response.Body)
	assert.Empty(t, e)
	return body
}
func SetRatingTest(t *testing.T, bookId int) {
	var baseResp books.BaseResponse
	randRating := rand.Intn(5)
	body := sendRequestAndGetBody(t, bookId, randRating)
	if err := json.Unmarshal(body, &baseResp); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, baseResp.Success, true)
}

func Test(t *testing.T) {
	randBook := getRandomBooks(t)
	t.Run("set rating test", func(t *testing.T) {
		SetRatingTest(t, randBook.BookId)
	})
}
