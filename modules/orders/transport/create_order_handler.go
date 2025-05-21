package ordertransport

import (
	itemstorage "dev-coffee-api/modules/items/storage"
	ordermodel "dev-coffee-api/modules/orders/model"
	orderservice "dev-coffee-api/modules/orders/service"
	orderstorage "dev-coffee-api/modules/orders/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data ordermodel.OrderCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := orderstorage.NewSQLStorage(db)
		itemStore := itemstorage.NewSQLStorage(db)
		service := orderservice.NewCreateOrderService(store, itemStore)

		err := service.CreateOrder(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": data.ID})
	}
}
