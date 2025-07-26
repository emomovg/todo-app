package repository

import (
	"context"
	"fmt"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/pkg/db"
)

type UserRepository struct {
	*db.Postgres
}

const TableName = "users"

func NewUserRepository(db *db.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) CreateUser(ctx context.Context, user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) values ($1, $2, $3) RETURNING id", TableName)
	row := u.Pool.QueryRow(ctx, query, user.UserName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
