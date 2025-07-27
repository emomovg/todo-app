package repository

import (
	"context"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/pkg/db"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUser(ctx context.Context, email, password string) (models.User, error)
}

type Todolist interface {
	Create(ctx context.Context, userId int, list models.TodoList) (int, error)
	GetAll(ctx context.Context, userId int) ([]models.TodoList, error)
	GetById(ctx context.Context, listId int) (models.TodoList, error)
	Delete(ctx context.Context, listId int) error
	Update(ctx context.Context, todoList models.TodoList) error
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
		Todolist:        NewTodolistRepository(db),
	}
}
