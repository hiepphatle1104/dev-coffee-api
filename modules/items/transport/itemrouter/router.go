package itemrouter

import (
	item "dev-coffee-api/modules/items/transport"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(v1 *gin.RouterGroup, db *gorm.DB) {
	router := v1.Group("/items")
	{
		router.GET("", item.GetItemsList(db))
		router.POST("", item.CreateNewItem(db))
		router.GET("/:id", item.GetItemByID(db))
		router.PATCH("/:id", item.UpdateItemByID(db))
		router.DELETE("/:id", item.DeleteItemByID(db))
	}
}
