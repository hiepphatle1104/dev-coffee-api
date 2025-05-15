package paymenttransport

import (
	paymentmodel "dev-coffee-api/modules/payments/model"
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

		var payments []paymentmodel.Payment
		if err := db.Find(&payments).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": payments, "paging": paging})
	}
}

func GetPaymentByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId, err := strconv.Atoi(c.Param("order-id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "order-id is required"})
			return
		}

		var payment paymentmodel.Payment
		if err = db.Where("order_id = ?", orderId).First(&payment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
