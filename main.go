package main

import (
	"dev-coffee-api/common"
	itemtransport "dev-coffee-api/modules/items/transport"
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

	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
