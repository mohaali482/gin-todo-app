package helpers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohaali482/todo/pkg/models"
	"gorm.io/gorm"
)

func GetObjectOr404(id string, db *gorm.DB, c *gin.Context) (models.Todo, bool) {
	var todo models.Todo
	// convert string id to int64 for preventing sql injection
	todoId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return models.Todo{}, false
	}

	result := db.First(&todo, "id = ?", todoId)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return models.Todo{}, false
	}

	return todo, true
}
