package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"UserId"`
}
type Auth interface {
	HashPassword(password string) (string, error)
	RegisterUser(firstName, lastName, username, password string) (*Token, error)
	Login(username, password string) string
	CheckPasswordHash(hashedPassword, password string) bool
	GenerateJwtToken(userId int) (string, error)
}
type Repository struct {
	Db                     *sql.DB
	SecretKey              string
	AccessTokenExpiration  int
	RefreshTokenExpiration int
}

func (s *Repository) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//RegisterUser by first name, last name, username, password
func (s *Repository) RegisterUser(firstName, lastName, username, password string) (*Token, error) {
	var LastInsertedId int
	hashedPassword, _ := s.HashPassword(password)
	_ = s.Db.QueryRow(`INSERT INTO "User" (first_name,last_name,username,password) VALUES ($1,$2,$3,$4)`,
		firstName, lastName, username, hashedPassword).Scan(&LastInsertedId)
	token, _ := s.GenerateAccessRefreshTokens(int(LastInsertedId))
	return token, nil

}

//Login by username and password
func (s *Repository) Login(username, password string) (*Token, error) {
	var tokens Token
	var UserPassword string
	var id int
	row := s.Db.QueryRow(`SELECT id, password FROM "User" as u WHERE u.username = $1`, username)
	switch err := row.Scan(&id, &UserPassword); err {
	case sql.ErrNoRows:
		return &tokens, sql.ErrNoRows
	case nil:
		if s.CheckPasswordHash(UserPassword, password) {
			token, exc := s.GenerateAccessRefreshTokens(id)
			if exc != nil {
				log.Fatal(fmt.Sprintf(`create tokens for user with UserId: %d err: %s`, id, exc.Error()))
				return &tokens, exc
			}
			return token, nil
		}
		return &tokens, errors.New(`password is invalid`)
	default:
		log.Fatal(fmt.Sprintf(`something was wrong err: %s`, err))
	}
	return &tokens, errors.New(`something was wrong`)
}
func (s *Repository) CheckPasswordHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func (s *Repository) GenerateAccessRefreshTokens(userId int) (*Token, error) {
	accessToken, err := s.generateJwtToken(userId, s.AccessTokenExpiration)
	if err != nil {
		log.Fatal(fmt.Sprintf(`create access token with user UserId:%d err: %s`, userId, err.Error()))
	}
	refreshToken, exc := s.generateJwtToken(userId, s.RefreshTokenExpiration)
	if exc != nil {
		log.Fatal(fmt.Sprintf(`create refresh token with user UserId:%d err: %s`, userId, err.Error()))
	}
	tokens := Token{accessToken, refreshToken}
	return &tokens, nil
}
func (s *Repository) generateJwtToken(userId, tokenExpiration int) (string, error) {
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(tokenExpiration) * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix()}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{standardClaims, userId}).SignedString([]byte(s.SecretKey))
}
