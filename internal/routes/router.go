package routes

import (
	"github.com/emomovg/todo-app/internal/services"
	"github.com/gin-gonic/gin"
)

type Router struct {
	service *services.Service
}

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", r.signUp)
		auth.POST("sign-in", r.signIn)
	}

	api := router.Group("api")
	{
		lists := api.Group("lists")
		{
			lists.POST("/", r.createList)
			lists.GET("/", r.getAllLists)
			lists.GET("/:id", r.getListById)
			lists.PUT("/:id", r.updateList)
			lists.DELETE("/:id", r.deleteList)

			items := api.Group(":id/items")
			{
				items.POST("", r.createItem)
				items.GET("", r.getAllItems)
				items.GET("/:item_id", r.getItemById)
				items.PUT("/:item_id", r.updateItem)
				items.DELETE("/:item_id", r.DeleteItem)

			}
		}
	}

	return router
}

func NewRouter(service *services.Service) *Router {
	return &Router{service: service}
}
