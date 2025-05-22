package paymentrouter

import (
	payment "dev-coffee-api/modules/payments/transport"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(v1 *gin.RouterGroup, db *gorm.DB) {
	router := v1.Group("/payments")
	{
		router.GET("", payment.GetPaymentsList(db))
		router.POST("", payment.CreateNewPayment(db))
		router.GET("/:order-id", payment.GetPaymentByID(db))
	}
}
