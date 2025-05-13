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

func UpdateDrink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		var data model.DrinkItemUpdate
		if err = c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStorage(db)
		biz := business.NewUpdateDrinkBusiness(store)

		err = biz.UpdateDrink(c.Request.Context(), id, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "updated",
			"data":    true,
		})

	}
}
