package ginDrink

import (
	"dev-coffee-api/modules/drinks/business"
	"dev-coffee-api/modules/drinks/model"
	"dev-coffee-api/modules/drinks/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ListDrink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging model.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Process()
		store := storage.NewSQLStorage(db)
		biz := business.NewListDrinkBusiness(store)

		if err := db.Table(model.DrinkItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		drinks, err := biz.ListDrink(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"paging": paging,
			"data":   drinks,
		})
	}
}
