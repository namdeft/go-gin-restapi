package transport

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/business"
	"gin-restapi/internal/dish/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteDish(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		store := storage.SQLStore(db)
		biz := business.DeleteDishBusiness(store)

		if err := biz.DeleteDish(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(true))
	}
}
