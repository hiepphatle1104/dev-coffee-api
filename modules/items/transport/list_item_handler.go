package itemtransport

import (
	"dev-coffee-api/common"
	itemmodel "dev-coffee-api/modules/items/model"
	itemservice "dev-coffee-api/modules/items/service"
	itemstorage "dev-coffee-api/modules/items/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetItemsList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}

		paging.Process()
		if err := db.Table(itemmodel.Item{}.TableName()).Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, common.NewBadRequestErrorResponse(err))
			return
		}

		store := itemstorage.NewSQLStorage(db)
		service := itemservice.NewGetItemListService(store)

		items, err := service.GetItemList(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewBadRequestErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponseWithPaging(items, &paging))
	}
}
