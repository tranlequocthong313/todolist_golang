package transport

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tranlequocthong313/todolist/module/item/business"
	"github.com/tranlequocthong313/todolist/module/item/helper"
	"github.com/tranlequocthong313/todolist/module/item/model"
	"github.com/tranlequocthong313/todolist/module/item/storage"
	"gorm.io/gorm"
)

func HandleCreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item model.TodoItem
		if err := c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item.Title = strings.TrimSpace(item.Title)
		storage := storage.NewMySQLStorage(db)
		business := business.NewCreateToDoItemBusiness(storage)
		if err := business.CreateNewItem(c.Request.Context(), &item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, helper.NewSimpleSuccessResponse(item.Id))
	}
}
