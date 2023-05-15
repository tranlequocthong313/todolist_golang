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

func HandleDeleteAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storage := storage.NewMySQLStorage(db)
		business := business.NewDeleteTodoItemBusiness(storage)
		if err := business.DeleteItem(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, helper.NewSimpleSuccessResponse(true))
	}
}
