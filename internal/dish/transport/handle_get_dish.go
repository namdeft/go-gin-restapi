package transport

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/business"
	"gin-restapi/internal/dish/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDish(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		store := storage.SQLStore(db)
		biz := business.GetDishBusiness(store)

		dish, err := biz.GetDish(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(dish))
	}
}
