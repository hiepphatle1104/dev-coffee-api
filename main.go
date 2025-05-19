package main

import (
	"dev-coffee-api/common"
	item "dev-coffee-api/modules/items/transport"
	order "dev-coffee-api/modules/orders/transport"
	payment "dev-coffee-api/modules/payments/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	port := common.EnvLookup("PORT")
	db := common.NewMySQLDatabase()

	router := gin.Default()

	v1 := router.Group("/api/v1")
	items := v1.Group("/items")
	{
		items.GET("", item.GetItemsList(db))
		items.POST("", item.CreateNewItem(db))
		items.GET("/:id", item.GetItemByID(db))
		items.PATCH("/:id", item.UpdateItemByID(db))
		items.DELETE("/:id", item.DeleteItemByID(db))
	}

	orders := v1.Group("/orders")
	{
		orders.GET("", order.GetOrdersList(db))
		orders.POST("", order.CreateOrder(db))
		orders.GET("/:id", order.GetOrderById(db))
		orders.PATCH("/:id", order.UpdateOrder(db))
		orders.DELETE("/:id", order.DeleteOrder(db))
	}

	payments := v1.Group("/payments")
	{
		payments.GET("", payment.GetPaymentsList(db))
		payments.POST("", payment.CreateNewPayment(db))
		payments.GET("/:order-id", payment.GetPaymentByID(db))
	}

	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
