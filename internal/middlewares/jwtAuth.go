package middlewares

import (
	"gin-restapi/internal/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
