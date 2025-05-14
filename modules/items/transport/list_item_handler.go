package itemtransport

import (
	itemmodel "dev-coffee-api/modules/items/model"
	itemservice "dev-coffee-api/modules/items/service"
	itemstorage "dev-coffee-api/modules/items/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetItemsList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging itemmodel.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()
		if err := db.Table(itemmodel.Item{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		store := itemstorage.NewSQLStorage(db)
		service := itemservice.NewGetItemListService(store)

		items, err := service.GetItemList(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": items, "paging": paging})
	}
}
