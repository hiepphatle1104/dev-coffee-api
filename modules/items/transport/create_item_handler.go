package itemtransport

import (
	"dev-coffee-api/common"
	itemmodel "dev-coffee-api/modules/items/model"
	itemservice "dev-coffee-api/modules/items/service"
	itemstorage "dev-coffee-api/modules/items/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateNewItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item itemmodel.ItemCreation
		if err := c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}

		store := itemstorage.NewSQLStorage(db)
		service := itemservice.NewCreateItemService(store)

		err := service.CreateNewItem(c.Request.Context(), &item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err))
			return
		}

		c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse(item.ID))
	}
}
