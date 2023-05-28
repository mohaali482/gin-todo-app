package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/todo/pkg/caches"
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
	pagination := helpers.Pagination{}
	limit, e := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if e != nil {
		limit = 10
	}
	page, e := strconv.Atoi(c.DefaultQuery("page", "1"))
	if e != nil {
		page = 1
	}
	pagination.Limit = limit
	pagination.Page = page

	// result := t.DB.Find(&todos)
	result := t.DB.Scopes(helpers.Paginate(todos, &pagination, t.DB)).Find(&todos)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	pagination.Items = todos
	c.JSON(http.StatusOK, pagination)
}

func (t *TodoController) GetTodo(c *gin.Context) {
	var todo models.Todo
	var found bool = false
	id := c.Param("id")
	err := caches.Get("todo:id:"+id, &todo)
	if err != nil {
		todo, found = helpers.GetObjectOr404(id, t.DB, c)
		if found {
			caches.Set("todo:id:"+id, todo)
		} else {
			return
		}
	} else {
		found = true
	}
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
	caches.Set("todo:id:"+strconv.Itoa(int(todo.ID)), todo)
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
	caches.Set("todo:id:"+id, todo)
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
	caches.Delete("todo:id:" + id)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
