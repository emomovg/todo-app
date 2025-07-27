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

func NewUserRepository(db *db.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) CreateUser(ctx context.Context, user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) values ($1, $2, $3) RETURNING id", models.TableName)
	row := u.Pool.QueryRow(ctx, query, user.UserName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserRepository) GetUser(ctx context.Context, username, passwordHash string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, username, email, password_hash FROM %s WHERE email = $1 AND password_hash = $2", models.TableName)
	row := u.Pool.QueryRow(ctx, query, username, passwordHash)

	err := row.Scan(
		&user.Id,
		&user.UserName,
		&user.Email,
		&user.Password,
	)

	return user, err
}
