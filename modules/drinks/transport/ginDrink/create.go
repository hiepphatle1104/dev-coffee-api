package ginDrink

import (
	"dev-coffee-api/modules/drinks/business"
	"dev-coffee-api/modules/drinks/model"
	"dev-coffee-api/modules/drinks/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateDrink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.DrinkItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

			return
		}

		store := storage.NewSQLStorage(db)
		biz := business.NewCreateDrinkBusiness(store)

		err := biz.CreateNewDrink(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "created",
			"data":    data.ID,
		})

	}
}
