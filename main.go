package main

import (
	"dev-coffee-api/modules/drinks/model"
	"dev-coffee-api/modules/drinks/transport/drinkRouter"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//type OrderStatus int
//
//const (
//	OrderStatusPending OrderStatus = iota
//	OrderStatusCompleted
//	OrderStatusCancelled
//)

type Order struct {
	ID           int         `json:"id" gorm:"column:id;"`
	CustomerName string      `json:"customer_name" gorm:"column:customer_name;"`
	Status       string      `json:"status" gorm:"column:status;"`
	OrderItems   []OrderItem `json:"order_items" gorm:"-"`
	CreatedAt    *time.Time  `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    *time.Time  `json:"updated_at" gorm:"column:updated_at;"`
}

func (Order) TableName() string { return "orders" }

type OrderItem struct {
	OrderID  int `json:"-" gorm:"column:order_id;"`
	ItemID   int `json:"item_id" gorm:"column:item_id;"`
	Quantity int `json:"quantity" gorm:"column:quantity;"`
}

func (OrderItem) TableName() string { return "order_items" }

type OrderCreation struct {
	ID           int         `json:"-" gorm:"column:id;"`
	CustomerName string      `json:"customer_name" gorm:"column:customer_name;"`
	OrderItems   []OrderItem `json:"order_items" gorm:"-"`
}

func (OrderCreation) TableName() string { return Order{}.TableName() }

type OrderUpdate struct {
	ID           int         `json:"-" gorm:"column:id;"`
	CustomerName string      `json:"customer_name" gorm:"column:customer_name;"`
	OrderItems   []OrderItem `json:"order_items" gorm:"-"`
	Status       string      `json:"status" gorm:"column:status;"`
}

func (OrderUpdate) TableName() string { return Order{}.TableName() }

type Payment struct {
	ID      int        `json:"id" gorm:"column:id;"`
	OrderID int        `json:"order_id" gorm:"column:order_id;"`
	Method  string     `json:"method" gorm:"column:method;"`
	PaidAt  *time.Time `json:"paid_at" gorm:"column:paid_at;"`
	Amount  float64    `json:"amount" gorm:"column:amount;"`
}

func (Payment) TableName() string { return "payments" }

type PaymentCreation struct {
	ID      int     `json:"-" gorm:"column:id;"`
	OrderID int     `json:"order_id" gorm:"column:order_id;"`
	Method  string  `json:"method" gorm:"column:method;"`
	Amount  float64 `json:"amount" gorm:"column:amount;"`
}

func (PaymentCreation) TableName() string { return Payment{}.TableName() }

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
		order.GET("/:id", GetOrderByID(db))
		order.PATCH("/:id", UpdateOrder(db))
		order.DELETE("/:id", DeleteOrder(db))
	}

	if err := router.Run(":" + port); err != nil {
		return
	}
}

func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data OrderUpdate
		if err = c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check item is valid
		for _, orderItem := range data.OrderItems {
			var item model.DrinkItem
			if err = db.Table(model.DrinkItem{}.TableName()).First(&item, orderItem.ItemID).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"data":  orderItem.ItemID,
				})

				return
			}
		}

		// Update
		if err = db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// TODO: Update item

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func DeleteOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = db.Table(Order{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func GetOrderByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data Order
		if err = db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data OrderCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Check valid item
		for _, orderItem := range data.OrderItems {
			var item model.DrinkItem
			if err := db.Table(model.DrinkItem{}.TableName()).First(&item, orderItem.ItemID).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"data":  orderItem.ItemID,
				})

				return
			}
		}

		// Create order
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Create Order_Items
		for _, orderItem := range data.OrderItems {
			orderItem.OrderID = data.ID
			if err := db.Create(&orderItem).Error; err != nil {

				db.Table(Order{}.TableName()).Where("id = ?", data.ID).Delete(nil)

				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})

				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data.ID,
		})
	}
}

func ListOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []Order
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
