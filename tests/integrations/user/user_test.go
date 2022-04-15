package user

import (
	"books/tests/integrations/utils"
	"testing"
)

func Test(t *testing.T) {
	utils.LoadEnv()
	const (
		userId       = 1
		newFirstName = "newFirstName"
		newLastName  = "newLastName"
		newUsername  = "newUsername"
		newPassword  = "newUnHashPassword"
	)
}
