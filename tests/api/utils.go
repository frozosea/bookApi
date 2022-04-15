package api_utils

import (
	"books/internal/auth"
	"books/internal/books"
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
	baseApiUrl = "http://localhost:8080"
)

func Login(username, password string) string {
	LoginUserStruct := auth.UserAuth{Username: username, Password: password}
	requestJson, err := json.Marshal(LoginUserStruct)
	var token auth.Token
	if err != nil {
		log.Fatal(err.Error())
	}
	response, exc := http.Post(fmt.Sprintf(`%s/auth`, baseApiUrl), "", bytes.NewBuffer(requestJson))
	if exc != nil {
		log.Fatal(exc.Error())
	}
	defer response.Body.Close()
	body, errs := io.ReadAll(response.Body)
	if errs != nil {
		log.Fatal(errs.Error())
	}
	if exc := json.Unmarshal(body, &token); exc != nil {
		log.Fatal(exc.Error())
	}
	return token.AccessToken
}
func getHttpClient() *http.Client {
	client := http.Client{}
	return &client
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
	assert.Equal(t, response.StatusCode, 200)
	return body
}
func CheckBaseResponse(t *testing.T, body []byte) {
	var result books.BaseResponse
	if exc := json.Unmarshal(body, &result); exc != nil {
		log.Fatalf(`wtf err:%s`, exc.Error())
	}
	assert.Equal(t, result.Success, true)
}
