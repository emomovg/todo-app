package routes

import (
	"github.com/emomovg/todo-app/internal/models"
	"github.com/emomovg/todo-app/internal/requests/todolist"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (r *Router) createList(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}

	var req models.TodoList
	if err := ctx.BindJSON(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.Service.Todolist.Create(ctx, userId.(int), req)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (r *Router) getAllLists(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}

	todolists, err := r.Service.Todolist.GetAll(ctx, userId.(int))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if len(todolists) == 0 {
		NewErrorResponse(ctx, http.StatusBadRequest, "it's id don't have todo lists")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"todo_lists": todolists,
	})
}

func (r *Router) getListById(ctx *gin.Context) {
	_, ok := ctx.Get("userId")
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	todolist, err := r.Service.Todolist.GetById(ctx, id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, todolist)
}

func (r *Router) updateList(ctx *gin.Context) {
	_, ok := ctx.Get("userId")
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	var req todolist.UpdateRequest

	if err := ctx.BindJSON(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tlist := req.ToTodoList()

	tlist.Id = id

	err = r.Service.Todolist.Update(ctx, *tlist)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})

}

func (r *Router) deleteList(ctx *gin.Context) {
	_, ok := ctx.Get("userId")
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	err = r.Service.Todolist.Delete(ctx, id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
