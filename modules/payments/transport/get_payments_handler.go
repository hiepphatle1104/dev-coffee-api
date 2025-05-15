package paymenttransport

import (
	paymentmodel "dev-coffee-api/modules/payments/model"
	paymentservice "dev-coffee-api/modules/payments/service"
	paymentstorage "dev-coffee-api/modules/payments/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetPaymentsList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging paymentmodel.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()

		store := paymentstorage.NewSQLStorage(db)
		service := paymentservice.NewGetPaymentsListService(store)

		data, err := service.GetPaymentsList(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data, "paging": paging})
	}
}

func GetPaymentByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId, err := strconv.Atoi(c.Param("order-id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "order-id is required"})
			return
		}

		store := paymentstorage.NewSQLStorage(db)
		service := paymentservice.NewGetPaymentByIDService(store)

		data, err := service.GetPaymentByID(c.Request.Context(), orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
