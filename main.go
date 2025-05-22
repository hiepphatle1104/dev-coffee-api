package main

import (
	"dev-coffee-api/common"
	"dev-coffee-api/modules/items/transport/itemrouter"
	"dev-coffee-api/modules/orders/transport/orderrouter"
	"dev-coffee-api/modules/payments/transport/paymentrouter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	common.LoadEnvFile()
	//gin.SetMode(gin.ReleaseMode)

	port := common.EnvLookup("PORT")
	appHost := common.EnvLookup("APP_HOST")
	db := common.NewMySQLDatabase()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{appHost}

	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	itemrouter.Initialize(v1, db)
	orderrouter.Initialize(v1, db)
	paymentrouter.Initialize(v1, db)

	log.Println("Server is running on port ", port)
	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
