package middlewares

import (
	"errors"
	"gin-restapi/internal/token"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := token.ExtractToken(c)
		jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid token",
				})
				c.Abort()
				return nil, errors.New("Invalid token")
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
			userID, exists := claims["user_id"].(float64)
			if !exists {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid userID in token",
				})
				c.Abort()
				return
			}

			c.Set("userID", int(userID))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
