package itemtransport

import (
	"dev-coffee-api/common"
	itemservice "dev-coffee-api/modules/items/service"
	itemstorage "dev-coffee-api/modules/items/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetItemByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}

		store := itemstorage.NewSQLStorage(db)
		service := itemservice.NewGetItemByIdService(store)

		item, err := service.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(item))
	}
}
