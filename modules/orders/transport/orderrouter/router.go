package orderrouter

import (
	order "dev-coffee-api/modules/orders/transport"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(v1 *gin.RouterGroup, db *gorm.DB) {
	router := v1.Group("/router")
	{
		router.GET("", order.GetOrdersList(db))
		router.POST("", order.CreateOrder(db))
		router.GET("/:id", order.GetOrderById(db))
		router.PATCH("/:id", order.UpdateOrder(db))
		router.DELETE("/:id", order.DeleteOrder(db))

		router.GET("/:id/items", order.GetOrderItems(db))
	}
}
