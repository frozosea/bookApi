package user_logger

import (
	logger "books/pkg/log"
	"github.com/sirupsen/logrus"
)

type IUserLogs interface {
	logger.IBaseExceptionLogger
	UpdateUsernameLog(userId int, newUsername string)
	UpdateFirstNameLog(userId int, newFirstName string)
	UpdateLastNameLog(userId int, newLastName string)
}
type UserLogger struct {
	Logger *logrus.Logger
}

func (u *UserLogger) ExceptionLog(err error) {
	u.Logger.Error(err.Error())
}
func (u *UserLogger) updateUserParamLog(userId int, updateValue string, newValue string) {
	u.Logger.Infof(`user id: %d new %s: %s`, userId, updateValue, newValue)
}
func (u *UserLogger) UpdateUsernameLog(userId int, newUsername string) {
	u.updateUserParamLog(userId, "username", newUsername)
}
func (u *UserLogger) UpdateFirstNameLog(userId int, newFirstName string) {
	u.updateUserParamLog(userId, "fist name", newFirstName)

}
func (u *UserLogger) UpdateLastNameLog(userId int, newLastName string) {
	u.updateUserParamLog(userId, "last name", newLastName)

}
