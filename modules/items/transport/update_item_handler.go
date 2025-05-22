package itemtransport

import (
	"dev-coffee-api/common"
	itemmodel "dev-coffee-api/modules/items/model"
	itemservice "dev-coffee-api/modules/items/service"
	itemstorage "dev-coffee-api/modules/items/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateItemByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}

		var item itemmodel.ItemUpdate
		if err = c.ShouldBind(&item); err != nil {
			c.JSON(http.StatusBadRequest, common.NewBadRequestErrorResponse(err))
			return
		}

		store := itemstorage.NewSQLStorage(db)
		service := itemservice.NewUpdateItemByIdService(store)

		err = service.UpdateItemById(c.Request.Context(), id, &item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(id))
	}
}
