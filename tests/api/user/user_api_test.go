package user_api_test

import (
	"books/internal/user"
	"books/pkg/models"
	api_utils "books/tests/api"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	username     = "testUsername"
	password     = "testPassword"
	baseApiUrl   = "http://localhost:8080"
	newFirstName = "newFirstNameApi"
	newLastName  = "newLastNameApi"
	newPassword  = "newPasswordApi"
)

func Test(t *testing.T) {
	accessToken := api_utils.Login(username, password)
	if accessToken == "" {
		accessToken = api_utils.Login(username, newPassword)
	}
	t.Run("Update Username test", func(t *testing.T) {
		newUsername := "UpdatedUsername"
		url := fmt.Sprintf(`%s/v1/user/UpdateUsername`, baseApiUrl)
		requestBody, err := json.Marshal(user.UpdateUsernameSchema{NewUsername: newUsername})
		assert.Empty(t, err)
		body := api_utils.SendRequestWithClient(t, "PUT", url, accessToken, requestBody)
		api_utils.CheckBaseResponse(t, body)
		api_utils.Login(newUsername, password)
	})
	t.Run("update first name test", func(t *testing.T) {
		url := fmt.Sprintf(`%s/v1/user/UpdateFirstName`, baseApiUrl)
		requestBody, err := json.Marshal(user.UpdateFirstNameSchema{NewFirstName: newFirstName})
		assert.Empty(t, err)
		body := api_utils.SendRequestWithClient(t, "PUT", url, accessToken, requestBody)
		api_utils.CheckBaseResponse(t, body)
	})
	t.Run("update last name test", func(t *testing.T) {
		url := fmt.Sprintf(`%s/v1/user/UpdateLastName`, baseApiUrl)
		scheme := user.UpdateLastNameSchema{NewLastName: newLastName}
		requestBody, err := json.Marshal(scheme)
		assert.Empty(t, err)
		body := api_utils.SendRequestWithClient(t, "PUT", url, accessToken, requestBody)
		api_utils.CheckBaseResponse(t, body)
	})
	t.Run("update password test", func(t *testing.T) {
		url := fmt.Sprintf(`%s/v1/user/UpdatePassword`, baseApiUrl)
		scheme := user.UpdatePasswordSchema{NewPassword: newPassword}
		requestBody, err := json.Marshal(scheme)
		assert.Empty(t, err)
		body := api_utils.SendRequestWithClient(t, "PUT", url, accessToken, requestBody)
		api_utils.CheckBaseResponse(t, body)
		api_utils.Login(username, newPassword)
	})
	t.Run("get ingo about user test", func(t *testing.T) {
		var responseSchema *models.User
		url := fmt.Sprintf(`%s/user/GetInfoAboutUser`, baseApiUrl)
		scheme := user.GetInfoAboutUserSchema{UserId: 1}
		requestBody, err := json.Marshal(scheme)
		assert.Empty(t, err)
		body := api_utils.SendRequestWithClient(t, "GET", url, "", requestBody)
		exc := json.Unmarshal(body, &responseSchema)
		assert.Empty(t, exc)
		assert.NotEmpty(t, responseSchema.Id)
		assert.NotEmpty(t, responseSchema.Username)
		assert.NotEmpty(t, responseSchema.FirstName)
		assert.NotEmpty(t, responseSchema.LastName)
	})
}
