package todolist

import (
	"github.com/emomovg/todo-app/internal/models"
	"strings"
)

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *UpdateRequest) ToTodoList() *models.TodoList {
	return &models.TodoList{
		Title:       strings.TrimSpace(r.Title),
		Description: r.Description,
	}
}
