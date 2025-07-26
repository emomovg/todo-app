package services

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/internal/repository"
	"os"
)

type AuthService struct {
	repo repository.IUserRepository
}

func NewAuthService(repo repository.IUserRepository) *AuthService {
	return &AuthService{repo}
}

func (a *AuthService) CreateUser(ctx context.Context, user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(ctx, user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("HASH_SALT"))))
}
