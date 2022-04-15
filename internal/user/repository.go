package user

import (
	"books/pkg/models"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type repo interface {
	UpdateUserParam
	GetUserInfo
}
type UpdateUserParam interface {
	UpdateFirstName(userId int, newFirstName string) (bool, error)
	UpdateLastName(userId int, newLastName string) (bool, error)
	UpdateUsername(userId int, newUsername string) (bool, error)
	UpdatePassword(userId int, newPassword string) (bool, error)
}

type GetUserInfo interface {
	GetInfoAboutUser(userId int) (models.User, error)
}

type Repository struct {
	Db *sql.DB
}

func (s *Repository) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *Repository) GetInfoAboutUser(userId int) (models.User, error) {
	var user models.User
	err := s.Db.QueryRow(`SELECT id,first_name,last_name,username FROM "User" as u  WHERE id = $1`, userId).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (s *Repository) updateUserParam(userId int, changeValue, newValue string) (bool, error) {
	var ExecRow = fmt.Sprintf(`UPDATE "User" AS u SET %s = '%s' WHERE id = %d`, changeValue, newValue, userId)
	_, exc := s.Db.Exec(ExecRow)
	if exc != nil {
		//log.Fatal(fmt.Sprintf(`update %s with user id: %d err: %s`, changeValue, userId, exc))
		return false, exc
	}
	return true, nil
}
func (s *Repository) UpdateFirstName(userId int, newFirstName string) (bool, error) {
	return s.updateUserParam(userId, `first_name`, newFirstName)
}
func (s *Repository) UpdateLastName(userId int, newLastName string) (bool, error) {
	return s.updateUserParam(userId, `last_name`, newLastName)
}
func (s *Repository) UpdateUsername(userId int, newUsername string) (bool, error) {
	return s.updateUserParam(userId, `username`, newUsername)
}

func (s *Repository) UpdatePassword(userId int, newPassword string) (bool, error) {
	hashPassword, _ := s.hashPassword(newPassword)
	return s.updateUserParam(userId, `password`, hashPassword)
}
