package paymenttransport

import (
	"dev-coffee-api/common"
	paymentservice "dev-coffee-api/modules/payments/service"
	paymentstorage "dev-coffee-api/modules/payments/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetPaymentsList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}
		paging.Process()

		store := paymentstorage.NewSQLStorage(db)
		service := paymentservice.NewGetPaymentsListService(store)

		data, err := service.GetPaymentsList(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponseWithPaging(data, &paging))
	}
}

func GetPaymentByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId, err := strconv.Atoi(c.Param("order-id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}

		store := paymentstorage.NewSQLStorage(db)
		service := paymentservice.NewGetPaymentByIDService(store)

		data, err := service.GetPaymentByID(c.Request.Context(), orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data))
	}
}
