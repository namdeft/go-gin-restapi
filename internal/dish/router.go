package dish

import (
	"gin-restapi/internal/database"
	"gin-restapi/internal/dish/controllers"
	"gin-restapi/internal/dish/storage"
	"gin-restapi/internal/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func RouterDish(r *gin.RouterGroup) {
	db, err := database.ConnectToMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	store := storage.SQLStore(db)

	dishes := r.Group("/dishes")
	dishes.Use(middlewares.JwtAuthMiddleware())
	{
		dishes.GET("/", controllers.GetDishes(store))
		dishes.GET("/:id", controllers.GetDish(store))
		dishes.POST("/", controllers.CreateDish(store))
		dishes.PUT("/:id", controllers.UpdateDish(store))
		dishes.DELETE("/:id", controllers.DeleteDish(store))
	}

}
