package ordertransport

import (
	"dev-coffee-api/common"
	itemstorage "dev-coffee-api/modules/items/storage"
	ordermodel "dev-coffee-api/modules/orders/model"
	orderservice "dev-coffee-api/modules/orders/service"
	orderstorage "dev-coffee-api/modules/orders/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}

		var data ordermodel.OrderUpdate
		if err = c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}

		store := orderstorage.NewSQLStorage(db)
		itemStore := itemstorage.NewSQLStorage(db)
		service := orderservice.NewUpdateOrderByIdService(store, itemStore)

		err = service.UpdateOrderByID(c.Request.Context(), id, &data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(id))
	}
}
