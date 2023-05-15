package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tranlequocthong313/todolist/module/item/business"
	"github.com/tranlequocthong313/todolist/module/item/helper"
	"github.com/tranlequocthong313/todolist/module/item/model"
	"github.com/tranlequocthong313/todolist/module/item/storage"
	"gorm.io/gorm"
)

func HandleListItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging model.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		storage := storage.NewMySQLStorage(db)
		business := business.NewListTodoItemBusiness(storage)
		paging.Process()
		items, err := business.ListItems(c, nil, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, helper.NewSuccessResponse(items, paging, nil))
	}
}
