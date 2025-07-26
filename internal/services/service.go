package services

import (
	"context"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/internal/repository"
)

type Todolist interface {
}

type UserService interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GenerateToken(tx context.Context, email, password string) (string, error)
}

type TodoItem interface {
}

type Service struct {
	UserService
	Todolist
	TodoItem
}

func NewService(userRepo repository.IUserRepository) *Service {
	return &Service{
		UserService: NewAuthService(userRepo),
	}
}
