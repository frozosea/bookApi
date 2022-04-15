package auth_api_test

import (
	"books/internal/auth"
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
	firstName = "testApiFirstName"
	lastName  = "testApiLastName"
	username  = "testApiUsername"
	password  = "testApiPassword"
)

func RegisterTest(t *testing.T, url string) {
	RegisterUserTestStruct := auth.UserRegister{Username: username, FirstName: firstName, Lastname: lastName, Password: password}
	requestJson, err := json.Marshal(RegisterUserTestStruct)
	var token auth.Token
	if err != nil {
		log.Fatal(err.Error())
	}
	response, err := http.Post(fmt.Sprintf(`%s/register`, url), "", bytes.NewBuffer(requestJson))
	defer response.Body.Close()
	body, errs := io.ReadAll(response.Body)
	if errs != nil {
		log.Fatal(errs.Error())
	}
	if exc := json.Unmarshal(body, &token); exc != nil {
		log.Fatal(exc.Error())
	}
	assert.Empty(t, err)
	assert.Equal(t, response.StatusCode, 201)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)

}
func LoginTest(t *testing.T, url string) {
	LoginUserStruct := auth.UserAuth{Username: username, Password: password}
	requestJson, err := json.Marshal(LoginUserStruct)
	var token auth.Token
	if err != nil {
		log.Fatal(err.Error())
	}
	response, err := http.Post(fmt.Sprintf(`%s/auth`, url), "", bytes.NewBuffer(requestJson))
	defer response.Body.Close()
	body, errs := io.ReadAll(response.Body)
	if errs != nil {
		log.Fatal(errs.Error())
	}
	if exc := json.Unmarshal(body, &token); exc != nil {
		log.Fatal(exc.Error())
	}
	assert.Empty(t, err)
	assert.Equal(t, response.StatusCode, 200)
	assert.NotEmpty(t, token.AccessToken)
	assert.NotEmpty(t, token.RefreshToken)
}

func Test(t *testing.T) {
	var url = "http://localhost:8080"
	t.Run("register and login test", func(t *testing.T) {
		RegisterTest(t, url)
		LoginTest(t, url)
	})
}
