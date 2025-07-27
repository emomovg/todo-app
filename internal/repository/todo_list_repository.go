package repository

import (
	"context"
	"fmt"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/pkg/db"
)

type TodolistRepository struct {
	*db.Postgres
}

func NewTodolistRepository(db *db.Postgres) *TodolistRepository {
	return &TodolistRepository{db}
}

func (tr *TodolistRepository) Create(ctx context.Context, userId int, list models.TodoList) (int, error) {
	tx, err := tr.Pool.Begin(ctx)
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", models.TodoListTable)
	row := tx.QueryRow(ctx, createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback(ctx)

		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", models.UserListTable)
	_, err = tx.Exec(ctx, createUsersListQuery, userId, id)

	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	return id, tx.Commit(ctx)
}

func (tr *TodolistRepository) GetAll(ctx context.Context, userId int) ([]models.TodoList, error) {
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		models.TodoListTable,
		models.UserListTable,
	)

	rows, err := tr.Pool.Query(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query todo lists: %w", err)
	}
	defer rows.Close()

	var lists []models.TodoList
	for rows.Next() {
		var list models.TodoList
		// Сканируем данные в структуру
		if err := rows.Scan(&list.Id, &list.Title, &list.Description); err != nil {
			return nil, fmt.Errorf("failed to scan todo list: %w", err)
		}
		lists = append(lists, list)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after rows iteration: %w", err)
	}

	return lists, nil
}

func (tr *TodolistRepository) GetById(ctx context.Context, listId int) (models.TodoList, error) {
	var todolist models.TodoList
	query := fmt.Sprintf("SELECT id, title, description FROM %s where id = $1", models.TodoListTable)

	row := tr.Pool.QueryRow(ctx, query, listId)
	err := row.Scan(
		&todolist.Id,
		&todolist.Title,
		&todolist.Description,
	)

	return todolist, err
}

func (tr *TodolistRepository) Delete(ctx context.Context, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s where id = $1", models.TodoListTable)

	_, err := tr.Pool.Exec(ctx, query, listId)

	return err
}

func (tr *TodolistRepository) Update(ctx context.Context, todoList models.TodoList) error {

}
