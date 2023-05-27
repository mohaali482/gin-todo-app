package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/todo/pkg/errors"
	"github.com/mohaali482/todo/pkg/helpers"
	"github.com/mohaali482/todo/pkg/models"
	"gorm.io/gorm"
)

type TodoController struct {
	DB *gorm.DB
}

func (t *TodoController) GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	result := t.DB.Find(&todos)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (t *TodoController) GetTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	todo, found := helpers.GetObjectOr404(id, t.DB, c)
	if !found {
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (t *TodoController) CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.ShouldBindJSON(&todo)
	err := todo.Validate()
	if err != nil {
		errors.ReturnErrorResponse(err, c)
		return
	}
	result := t.DB.Create(&todo)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (t *TodoController) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	todo, found := helpers.GetObjectOr404(id, t.DB, c)
	if !found {
		return
	}
	c.BindJSON(&todo)
	result := t.DB.Save(&todo)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return

	}
	c.JSON(http.StatusOK, todo)
}

func (t *TodoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	todo, found := helpers.GetObjectOr404(id, t.DB, c)
	if !found {
		return
	}
	result := t.DB.Delete(&todo)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
