package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tranlequocthong313/todolist/module/item/business"
	"github.com/tranlequocthong313/todolist/module/item/helper"
	"github.com/tranlequocthong313/todolist/module/item/model"
	"github.com/tranlequocthong313/todolist/module/item/storage"
	"gorm.io/gorm"
)

func HandleUpdateAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var item model.TodoItem
		if err := c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storage := storage.NewMySQLStorage(db)
		business := business.NewUpdateTodoItemBusiness(storage)
		if err := business.UpdateItem(c, map[string]interface{}{"id": id}, &item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, helper.NewSimpleSuccessResponse(true))
	}
}
