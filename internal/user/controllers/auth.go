package controllers

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/token"
	"gin-restapi/internal/user/dto"
	"gin-restapi/internal/user/services"
	"gin-restapi/internal/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(store *storage.SqlStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.LoginInput

		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vErr := input.ValidateLoginInput()
		if vErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": vErr.Error(),
			})

			return
		}

		biz := services.AuthService(store)

		tokenKey, err := biz.Login(c.Request.Context(), &input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token.ExtractToken(c)

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(tokenKey))

	}
}

func Register(store *storage.SqlStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.RegisterInput

		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vErr := input.ValidateRegisterInput()
		if vErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": vErr.Error(),
			})

			return
		}

		biz := services.AuthService(store)

		if err := biz.Register(c.Request.Context(), &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(input))
	}
}
