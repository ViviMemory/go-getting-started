package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/pkg/repository"
	"time"
)

const (
	salt       = "sdfsdfweqe34kl23423mkpxz"
	signingKey = "23j3ji0frfe90snjcmlas3"
	tokenTTL   = 36600 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repo repository.Authentication) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CheckAuth(phone string) (int, error) {
	id, err := s.repo.CheckAuth(phone)
	return id, err
}

func (s *AuthService) CreateUser(user model.SignUpInput) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(name, phone string) (string, error) {
	user, err := s.repo.GetUser(name, phone)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))

}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *token claims")
	}

	return claims.UserId, nil
}
