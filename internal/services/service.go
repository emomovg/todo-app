package services

import (
	"context"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/internal/repository"
)

type Todolist interface {
	Create(ctx context.Context, userId int, list models.TodoList) (int, error)
	GetAll(ctx context.Context, userId int) ([]models.TodoList, error)
	GetById(ctx context.Context, listId int) (models.TodoList, error)
	Delete(ctx context.Context, listId int) error
	Update(ctx context.Context, todolist models.TodoList) error
}

type UserService interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GenerateToken(ctx context.Context, email, password string) (string, error)
	ParseToken(ctx context.Context, token string) (int, error)
}

type TodoItem interface {
}

type Service struct {
	UserService
	Todolist
	TodoItem
}

func NewService(userRepo repository.IUserRepository, tlRepo repository.Todolist) *Service {
	return &Service{
		UserService: NewAuthService(userRepo),
		Todolist:    NewTodolistService(tlRepo),
	}
}
