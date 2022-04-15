package user

import (
	user_logger "books/pkg/log/user"
	"books/pkg/models"
)

type ServiceInterface interface {
	UpdateFirstName(userId int, newFirstName string) (bool, error)
	UpdateLastName(userId int, newLastName string) (bool, error)
	UpdateUsername(userId int, newUsername string) (bool, error)
	UpdatePassword(userId int, newPassword string) (bool, error)
	GetInfoAboutUser(userId int) (models.User, error)
}
type Service struct {
	Repository repo
	SecretKey  string
	Logger     user_logger.IUserLogs
}

func (s *Service) UpdateFirstName(userId int, newFirstName string) (bool, error) {
	ok, err := s.Repository.UpdateFirstName(userId, newFirstName)
	if err != nil {
		go s.Logger.ExceptionLog(err)
		return ok, err
	}
	go s.Logger.UpdateFirstNameLog(userId, newFirstName)
	return ok, err
}
func (s *Service) UpdateLastName(userId int, newLastName string) (bool, error) {
	ok, err := s.Repository.UpdateLastName(userId, newLastName)
	if err != nil {
		go s.Logger.ExceptionLog(err)
		return ok, err
	}
	go s.Logger.UpdateLastNameLog(userId, newLastName)
	return ok, err
}
func (s *Service) UpdateUsername(userId int, newUsername string) (bool, error) {
	ok, err := s.Repository.UpdateUsername(userId, newUsername)
	if err != nil {
		go s.Logger.ExceptionLog(err)
		return ok, err
	}
	go s.Logger.UpdateUsernameLog(userId, newUsername)
	return ok, err
}
func (s *Service) UpdatePassword(userId int, newPassword string) (bool, error) {
	ok, err := s.Repository.UpdatePassword(userId, newPassword)
	if err != nil {
		go s.Logger.ExceptionLog(err)
		return ok, err
	}
	return ok, err
}
func (s *Service) GetInfoAboutUser(userId int) (models.User, error) {
	user, err := s.Repository.GetInfoAboutUser(userId)
	if err != nil {
		s.Logger.ExceptionLog(err)
		return user, err
	}
	return user, err
}
