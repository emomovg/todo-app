package services

import (
	"context"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/internal/repository"
)

type TodolistService struct {
	repo repository.Todolist
}

func NewTodolistService(repo repository.Todolist) *TodolistService {
	return &TodolistService{repo}
}

func (ts *TodolistService) Create(ctx context.Context, userId int, list models.TodoList) (int, error) {
	return ts.repo.Create(ctx, userId, list)
}

func (ts *TodolistService) GetAll(ctx context.Context, userId int) ([]models.TodoList, error) {
	return ts.repo.GetAll(ctx, userId)
}

func (ts *TodolistService) GetById(ctx context.Context, listId int) (models.TodoList, error) {
	return ts.repo.GetById(ctx, listId)
}

func (ts *TodolistService) Delete(ctx context.Context, listId int) error {
	return ts.repo.Delete(ctx, listId)
}

func (ts *TodolistService) Update(ctx context.Context, todoList models.TodoList) error {
	return ts.repo.Update(ctx, todoList)
}
