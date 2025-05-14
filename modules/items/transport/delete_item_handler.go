package itemtransport

import (
	itemservice "dev-coffee-api/modules/items/service"
	itemstorage "dev-coffee-api/modules/items/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteItemByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}

		store := itemstorage.NewSQLStorage(db)
		service := itemservice.NewDeleteItemByIdService(store)

		err = service.DeleteItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
