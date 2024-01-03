package transport

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/business"
	"gin-restapi/internal/dish/model"
	"gin-restapi/internal/dish/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDish(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.DishCreation

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.SQLStore(db)
		biz := business.CreateDishBusiness(store)

		if err := biz.CreateNewDish(c.Request.Context(), &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(input.Id))
	}
}
