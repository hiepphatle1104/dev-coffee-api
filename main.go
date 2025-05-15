package main

import (
	"dev-coffee-api/common"
	itemtransport "dev-coffee-api/modules/items/transport"
	ordertransport "dev-coffee-api/modules/orders/transport"
	paymenttransport "dev-coffee-api/modules/payments/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	port := common.EnvLookup("PORT")
	db := common.NewMySQLDatabase()

	router := gin.Default()

	v1 := router.Group("/api/v1")
	items := v1.Group("/items")
	{
		items.GET("", itemtransport.GetItemsList(db))
		items.POST("", itemtransport.CreateNewItem(db))
		items.GET("/:id", itemtransport.GetItemByID(db))
		items.PATCH("/:id", itemtransport.UpdateItemByID(db))
		items.DELETE("/:id", itemtransport.DeleteItemByID(db))
	}

	orders := v1.Group("/orders")
	{
		orders.GET("", ordertransport.GetOrdersList(db))
		orders.POST("", ordertransport.CreateOrder(db))
		orders.GET("/:id", ordertransport.GetOrderById(db))
		orders.PATCH("/:id", ordertransport.UpdateOrder(db))
		orders.DELETE("/:id", ordertransport.DeleteOrder(db))
	}

	payments := v1.Group("/payments")
	{
		payments.GET("", paymenttransport.GetPaymentsList(db))
		payments.POST("", paymenttransport.CreateNewPayment(db))
		payments.GET("/:order-id", paymenttransport.GetPaymentByID(db))
	}

	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
