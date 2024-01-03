package dish

import (
	"gin-restapi/internal/database"
	"gin-restapi/internal/dish/transport"
	"gin-restapi/internal/middlewares"

	"log"

	"github.com/gin-gonic/gin"
)

func RouterDish(r *gin.RouterGroup) {
	db, err := database.ConnectToMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	dishes := r.Group("/dishes")
	dishes.Use(middlewares.JwtAuthMiddleware())
	{
		dishes.GET("/", transport.GetDishes(db))
		dishes.GET("/:id", transport.GetDish(db))
		dishes.POST("/", transport.CreateDish(db))
		dishes.PUT("/:id", transport.UpdateDish(db))
		dishes.DELETE("/:id", transport.DeleteDish(db))
	}

}
