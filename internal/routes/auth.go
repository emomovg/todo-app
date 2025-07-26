package routes

import (
	"github.com/emomovg/todo-app/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) signUp(ctx *gin.Context) {
	var input models.User

	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := r.service.UserService.CreateUser(ctx, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (r *Router) signIn(ctx *gin.Context) {

}
