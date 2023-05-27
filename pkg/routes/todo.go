package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohaali482/todo/pkg/controllers"
	"gorm.io/gorm"
)

func TodoRoutes(r *gin.Engine, db *gorm.DB) {
	t := controllers.TodoController{
		DB: db,
	}

	routes := r.Group("/todos")
	// Data Retrieval
	routes.GET("/", t.GetAllTodos)
	routes.GET("/:id", t.GetTodo)
	// Data Manipulation
	routes.POST("/", t.CreateTodo)
	routes.PUT("/:id", t.UpdateTodo)
	routes.DELETE("/:id", t.DeleteTodo)
}
