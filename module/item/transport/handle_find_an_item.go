package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tranlequocthong313/todolist/module/item/business"
	"github.com/tranlequocthong313/todolist/module/item/helper"
	"github.com/tranlequocthong313/todolist/module/item/storage"
	"gorm.io/gorm"
)

func HandleFindAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storage := storage.NewMySQLStorage(db)
		business := business.NewFindTodoItemBusiness(storage)
		item, err := business.FindAnItem(c, map[string]interface{}{"id": id})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, helper.NewSimpleSuccessResponse(item))
	}
}
