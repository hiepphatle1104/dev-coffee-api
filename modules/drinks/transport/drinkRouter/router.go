package drinkRouter

import (
	"dev-coffee-api/modules/drinks/transport/ginDrink"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("", ginDrink.CreateDrink(db))
	router.GET("", ginDrink.ListDrink(db))
	router.GET("/:id", ginDrink.GetDrink(db))
	router.PATCH("/:id", ginDrink.UpdateDrink(db))
	router.DELETE("/:id", ginDrink.DeleteDrink(db))
}
