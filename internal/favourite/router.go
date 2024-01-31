package favourite

import (
	"gin-restapi/internal/favourite/controllers"
	"gin-restapi/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RouterFavourite(r *gin.RouterGroup, controller controllers.FavouriteController) {
	favourites := r.Group("/favourites")
	favourites.Use(middlewares.JwtAuthMiddleware())
	{
		favourites.GET("/", controller.GetFavourites())
		favourites.POST("/:dish_id", controller.AddFavourite())
		favourites.DELETE("/:dish_id", controller.DeleteFavourite())
	}
}
