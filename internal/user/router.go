package user

import (
	"gin-restapi/internal/user/controllers"

	"github.com/gin-gonic/gin"
)

func RouterAuth(r *gin.RouterGroup, controller controllers.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controller.Register())
		auth.POST("/login", controller.Login())
	}

}
