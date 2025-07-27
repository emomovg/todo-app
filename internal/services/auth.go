package services

import (
	"context"
	"crypto/sha1"
	"errors"
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

type jwtClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: user.Id,
	})

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func (a *AuthService) ParseToken(ctx context.Context, accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing Method")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *jwtClaims")
	}

	return claims.UserID, nil
}
