package user

import (
	"github.com/emomovg/todo-app/internal/models"
	"strings"
)

type AuthRequest struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *AuthRequest) ToUser() *models.User {
	return &models.User{
		UserName: strings.TrimSpace(r.UserName),
		Email:    strings.ToLower(r.Email),
		Password: r.Password,
	}
}
