package transport

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/business"
	"gin-restapi/internal/dish/model"
	"gin-restapi/internal/dish/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateDish(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.DishUpdation

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		store := storage.SQLStore(db)
		biz := business.UpdateDishBusiness(store)

		if err := biz.UpdateDish(c.Request.Context(), id, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(true))
	}
}
