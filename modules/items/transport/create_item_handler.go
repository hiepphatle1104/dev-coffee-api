package itemtransport

import (
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := itemstorage.NewSQLStorage(db)
		service := itemservice.NewCreateItemService(store)

		err := service.CreateNewItem(c.Request.Context(), &item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": true})

	}
}
