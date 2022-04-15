package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type ServiceAuth interface {
	RegisterUser(firstName, lastName, username, password string) (string, error)
	Login(username, password string) (string, error)
	Refresh(token string) (string, error)
}
type Service struct {
	Repo      *Repository
	SecretKey string
}

func (s *Service) RegisterUser(firstName, lastName, username, password string) (*Token, error) {
	return s.Repo.RegisterUser(firstName, lastName, username, password)
}
func (s *Service) Login(username, password string) (*Token, error) {
	return s.Repo.Login(username, password)
}

func GetUserIdByJwtToken(accessToken, secretKey string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New(`token claims are not valid`)
	}
	return claims.UserId, nil
}
