package ordertransport

import (
	ordermodel "dev-coffee-api/modules/orders/model"
	orderservice "dev-coffee-api/modules/orders/service"
	orderstorage "dev-coffee-api/modules/orders/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetOrdersList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging ordermodel.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()
		if err := db.Table(ordermodel.Order{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		store := orderstorage.NewSQLStorage(db)
		service := orderservice.NewGetOrdersService(store)

		data, err := service.GetOrders(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data, "paging": paging})
	}
}

func GetOrderById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}

		store := orderstorage.NewSQLStorage(db)
		service := orderservice.NewGetOrderByIdService(store)

		data, err := service.GetOrderById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
