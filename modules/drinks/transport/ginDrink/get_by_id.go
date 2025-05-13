package ginDrink

import (
	"dev-coffee-api/modules/drinks/business"
	"dev-coffee-api/modules/drinks/model"
	"dev-coffee-api/modules/drinks/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetDrink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStorage(db)
		biz := business.NewGetDrinkBusiness(store)
		var result *model.DrinkItem

		result, err = biz.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
