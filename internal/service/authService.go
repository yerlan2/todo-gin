package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	"yn/todo/internal/model"
	"yn/todo/internal/repository"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt      = "djhfjsdkhsjdfjdfhdsjkfhshfsjfhdsjfhkjhuehuwhehcneu"
	signinKey = "hdfdshflakhfdkfjdahskfhdksfhdkfhsdhfkjdshaldkfhdak"
	tokenTTL  = 12 * time.Hour
)

// Field
type AuthService struct {
	repo repository.Authorization
}
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// Constructor
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Function
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

// Method
func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
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
	return token.SignedString([]byte(signinKey))
}
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return []byte(signinKey), nil
		},
	)
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}
