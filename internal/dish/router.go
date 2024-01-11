package dish

import (
	"gin-restapi/internal/dish/controllers"
	"gin-restapi/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RouterDish(r *gin.RouterGroup, controller controllers.DishController) {
	dishes := r.Group("/dishes")
	dishes.Use(middlewares.JwtAuthMiddleware())
	{
		dishes.GET("/", controller.GetDishes())
		dishes.GET("/:id", controller.GetDish())
		dishes.POST("/", controller.CreateDish())
		dishes.PUT("/:id", controller.UpdateDish())
		dishes.DELETE("/:id", controller.DeleteDish())
	}

}
