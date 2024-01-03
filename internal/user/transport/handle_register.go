package transport

import (
	// "fmt"
	"gin-restapi/internal/common"
	"gin-restapi/internal/user/business"
	"gin-restapi/internal/user/model"
	"gin-restapi/internal/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.RegisterInput

		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.SQLStore(db)
		biz := business.RegisterBusiness(store)

		if err := biz.Register(c.Request.Context(), &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(input))
	}
}
