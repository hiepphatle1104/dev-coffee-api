package paymenttransport

import (
	ordermodel "dev-coffee-api/modules/orders/model"
	paymentmodel "dev-coffee-api/modules/payments/model"
	"fmt"
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

		var exists paymentmodel.Payment
		if err := db.Where("order_id = ?", data.OrderID).First(&exists).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "some error while creating payment"})
			return
		}

		// Summarize total amount
		var amount float64 = 0
		var orderItems []ordermodel.OrderItem
		if err := db.Table(ordermodel.OrderItem{}.TableName()).Where("order_id = ?", data.OrderID).Find(&orderItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, item := range orderItems {
			amount += item.UnitPrice * float64(item.Quantity)
		}

		data.Amount = amount

		fmt.Println(data)

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
