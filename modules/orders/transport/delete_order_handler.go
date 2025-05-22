package ordertransport

import (
	"dev-coffee-api/common"
	orderservice "dev-coffee-api/modules/orders/service"
	orderstorage "dev-coffee-api/modules/orders/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}

		store := orderstorage.NewSQLStorage(db)
		service := orderservice.NewDeleteOrderByIdService(store)

		err = service.DeleteOrderByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
