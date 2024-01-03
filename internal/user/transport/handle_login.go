package transport

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/token"
	"gin-restapi/internal/user/business"
	"gin-restapi/internal/user/model"
	"gin-restapi/internal/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.LoginInput

		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.SQLStore(db)
		biz := business.LoginBusiness(store)

		tokenKey, err := biz.Login(c.Request.Context(), &input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token.ExtractToken(c)

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(tokenKey))

	}
}
