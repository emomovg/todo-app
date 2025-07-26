package repository

import (
	"context"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/pkg/db"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
}

type Todolist interface {
}

type TodoItem interface {
}

type Repository struct {
	IUserRepository
	Todolist
	TodoItem
}

func NewRepository(db *db.Postgres) *Repository {
	return &Repository{
		IUserRepository: NewUserRepository(db),
	}
}
