package ginDrink

import (
	"dev-coffee-api/modules/drinks/business"
	"dev-coffee-api/modules/drinks/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteDrink(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStorage(db)
		biz := business.NewDeleteDrinkBusiness(store)

		if err = biz.DeleteDrink(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "deleted",
			"data":    true,
		})
	}
}
