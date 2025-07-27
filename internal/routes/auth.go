package routes

import (
	"fmt"
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/internal/requests/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) signUp(ctx *gin.Context) {
	var req user.AuthRequest
	var user *models.User

	if err := ctx.BindJSON(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user = req.ToUser()

	id, err := r.Service.UserService.CreateUser(ctx, *user)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(id)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (r *Router) signIn(ctx *gin.Context) {
	var req user.LoginRequest

	if err := ctx.BindJSON(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	token, err := r.Service.UserService.GenerateToken(ctx, req.Email, req.Password)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
