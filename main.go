package main

import (
	"dev-coffee-api/common"
	"dev-coffee-api/modules/items/transport/itemrouter"
	"dev-coffee-api/modules/orders/transport/orderrouter"
	"dev-coffee-api/modules/payments/transport/paymentrouter"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	// Static route
	router.Static("/uploads", "./uploads")
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
			return
		}

		// Create a folder if not exists
		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			return
		}

		// Rename file
		ext := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + ext

		// Save path
		filePath := filepath.Join("uploads", newFileName)
		if err = c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
			return
		}

		// Response
		imageURL := fmt.Sprintf("http://localhost:%s/%s", port, filePath)
		c.JSON(http.StatusOK, common.NewSuccessResponse(imageURL))
	})

	log.Println("Server is running on port ", port)
	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
