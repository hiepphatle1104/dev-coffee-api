package main

import (
	"dev-coffee-api/modules/drinks/transport/drinkRouter"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

//type OrderStatus int
//
//const (
//	OrderStatusPending OrderStatus = iota
//	OrderStatusCompleted
//	OrderStatusCancelled
//)

type Orders struct {
	ID           int    `json:"id" gorm:"column:id;"`
	CustomerName string `json:"customer_name" gorm:"column:customer_name;"`
	TotalPrice   int    `json:"total_price" gorm:"column:total_price;"`
	Status       string `json:"status" gorm:"column:status;"`
	CreatedAt    string `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    string `json:"updated_at" gorm:"column:updated_at;"`
}

type OrderItem struct {
}

type OrderCreation struct {
	ID           int    `json:"-" gorm:"column:id;"`
	CustomerName string `json:"customer_name" gorm:"column:customer_name;"`
}

func main() {
	dsn, ok := os.LookupEnv("DB_CONN_URL")
	if !ok {
		log.Fatalln("DB_CONN_URL not found")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatalln("PORT not found")
	}

	// Connect to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	router := gin.Default()
	v1 := router.Group("/api/v1")
	drinkRouter.Router(v1.Group("/drinks"), db)

	order := v1.Group("/orders")
	{
		order.GET("", ListOrders(db))
		order.POST("", CreateOrder(db))
		order.GET("/:id")
		order.PATCH("/:id")
		order.DELETE("/:id")
	}

	if err := router.Run(":" + port); err != nil {
		return
	}
}

func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order Orders

		if err := c.ShouldBind(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Create(&order).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": order.ID,
		})
	}
}

func ListOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []Orders

		if err := db.Find(&orders).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": orders,
		})
	}
}
