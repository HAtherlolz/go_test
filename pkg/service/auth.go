package service

import (
	"fmt"
	"crypto/sha1"
	"github.com/Hatherlolz/go_test/pkg/repository"
	"github.com/Hatherlolz/go_test"
)

const salt = "sdadad3241sa"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
    return &AuthService{
        repo: repo,
    }
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}