package services

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
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

func (a *AuthService) GenerateToken(ctx context.Context, email, password string) (string, error) {
	passwordHash := generatePasswordHash(password)
	user, err := a.repo.GetUser(ctx, email, passwordHash)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.Id,
		"expireAt": time.Now().Add(time.Hour * 24).Unix(),
		"issueAt":  time.Now().Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
