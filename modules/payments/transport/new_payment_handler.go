package paymenttransport

import (
	orderstorage "dev-coffee-api/modules/orders/storage"
	paymentmodel "dev-coffee-api/modules/payments/model"
	paymentservice "dev-coffee-api/modules/payments/service"
	paymentstorage "dev-coffee-api/modules/payments/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateNewPayment(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data paymentmodel.PaymentCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := paymentstorage.NewSQLStorage(db)
		orderItemStore := orderstorage.NewSQLStorage(db)
		orderStore := orderstorage.NewSQLStorage(db)
		service := paymentservice.NewCreateNewPaymentService(store, orderItemStore, orderStore)

		err := service.CreateNewPayment(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
