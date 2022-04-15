package authTest

import (
	"books/internal/auth"
	"books/settings"
	"books/tests/integrations/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const (
	firstName = "testFirstName"
	lastName  = "testLastName"
	username  = "testUsername"
	password  = "testPassword"
)

func Test(t *testing.T) {
	utils.LoadEnv()
	db, err := settings.GetDatabase()
	if err != nil {
		log.Fatalf(`can not connect to database err: %s`, err)
	}
	repo := auth.Repository{
		Db:                     db,
		SecretKey:              "testSecretKey",
		AccessTokenExpiration:  12,
		RefreshTokenExpiration: 24}
	t.Run("register", func(t *testing.T) {
		result, err := repo.RegisterUser(firstName, lastName, username, password)
		errOk := assert.Empty(t, err)
		if !errOk {
			t.Error(`assertion error`)
		}
		tokensOk := assert.NotEmpty(t, result.AccessToken, result.RefreshToken)
		if !tokensOk {
			t.Error(`assertion error`)
		}
	})
	t.Run("test login", func(t *testing.T) {
		result, err := repo.Login(username, password)
		errOk := assert.Empty(t, err)
		if !errOk {
			t.Error(`assertion error`)
		}
		tokenOk := assert.NotEmpty(t, result.AccessToken, result.RefreshToken)
		if !tokenOk {
			t.Error(`assertion error`)
		}
	})
}
