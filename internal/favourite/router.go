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
		favourites.GET("/:user_id", controller.GetFavourites())
		favourites.POST("/:user_id/:dish_id", controller.AddFavourite())
		favourites.DELETE("/:user_id/:dish_id", controller.DeleteFavourite())
	}
}
