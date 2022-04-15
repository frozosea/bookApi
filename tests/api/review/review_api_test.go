package review_api_test

import (
	"books/internal/books"
	"books/internal/review"
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
	username  = "testApiUsername"
	password  = "testApiPassword"
	Review    = "this is test review for api"
	newReview = "new review for api"
	baseUrl   = "http://localhost:8080"
	BookId    = 1
)

func getClient() *http.Client {
	return &http.Client{}
}
func sendRequestWithMethodAndJson(t *testing.T, method, accessToken, url string, requestBody []byte) []byte {
	client := getClient()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	assert.Empty(t, err)
	req.Header.Set(`Authorization`, fmt.Sprintf(`Bearer %s`, accessToken))
	response, exc := client.Do(req)
	assert.Empty(t, exc)
	defer response.Body.Close()
	body, errs := io.ReadAll(response.Body)
	assert.Empty(t, errs)
	return body
}
func assertReviewsIsNotEmpty(t *testing.T, review *models.UserReview) {
	assert.NotEmpty(t, review.Id)
	assert.NotEmpty(t, review.Review)
	assert.NotEmpty(t, review.ReviewId)
	assert.NotEmpty(t, review.Username)
	assert.NotEmpty(t, review.FirstName)
	assert.NotEmpty(t, review.LastName)
}
func WriteReviewTest(t *testing.T, accessToken string, bookId int) int {
	requestSchema := review.WriteReviewSchema{BookID: bookId, Review: Review}
	requestBody, exc := json.Marshal(requestSchema)
	assert.Empty(t, exc)
	writeReviewUrl := fmt.Sprintf(`%s/v1/review/WriteReview`, baseUrl)
	var reviewResponse review.WriteReviewResponseSchema
	body := sendRequestWithMethodAndJson(t, "POST", accessToken, writeReviewUrl, requestBody)
	if err := json.Unmarshal(body, &reviewResponse); err != nil {
		log.Fatal(err.Error())
	}
	assert.NotEmpty(t, reviewResponse.ReviewId)
	return reviewResponse.ReviewId
}
func GetAllReviewsByBookIdTest(t *testing.T, bookId int) {
	requestSchema := review.BookId{BookId: bookId}
	var responseSchema []*models.UserReview
	requestBody, exc := json.Marshal(requestSchema)
	assert.Empty(t, exc)
	url := fmt.Sprintf(`%s/review/GetReviewsByBookId`, baseUrl)
	body := sendRequestWithMethodAndJson(t, "GET", "", url, requestBody)
	if err := json.Unmarshal(body, &responseSchema); err != nil {
		log.Fatal(err.Error())
	}
	for _, book := range responseSchema {
		assertReviewsIsNotEmpty(t, book)
	}
}
func UpdateReviewTest(t *testing.T, accessToken string, reviewId int) {
	url := fmt.Sprintf(`%s/v1/review/UpdateReview`, baseUrl)
	requestSchema := review.UpdateReviewSchema{ReviewId: reviewId, NewReview: newReview}
	requestBody, exc := json.Marshal(requestSchema)
	var responseBody *books.BaseResponse
	assert.Empty(t, exc)
	body := sendRequestWithMethodAndJson(t, "PUT", accessToken, url, requestBody)
	if err := json.Unmarshal(body, &responseBody); err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, responseBody.Success, true)
}
func DeleteReviewTest(t *testing.T, accessToken string, reviewId int) {
	url := fmt.Sprintf(`%s/v1/review/DeleteReview`, baseUrl)
	requestSchema := review.DeleteReviewSchema{ReviewId: reviewId}
	requestBody, exc := json.Marshal(requestSchema)
	var responseBody *books.BaseResponse
	assert.Empty(t, exc)
	body := sendRequestWithMethodAndJson(t, "DELETE", accessToken, url, requestBody)
	if err := json.Unmarshal(body, &responseBody); err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, responseBody.Success, true)
}
func Test(t *testing.T) {
	accessToken := api_utils.Login(username, password)
	t.Run("test review api", func(t *testing.T) {
		reviewId := WriteReviewTest(t, accessToken, BookId)
		UpdateReviewTest(t, accessToken, reviewId)
		GetAllReviewsByBookIdTest(t, BookId)
		DeleteReviewTest(t, accessToken, reviewId)
	})
}
