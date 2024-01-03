package transport

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/business"
	"gin-restapi/internal/dish/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDishes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Process()

		store := storage.SQLStore(db)
		biz := business.GetDishesBusiness(store)

		dishes, err := biz.GetDishes(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusAccepted, common.SuccessResponse(dishes, paging, nil))
	}
}
