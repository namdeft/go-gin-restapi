package user

import (
	"gin-restapi/internal/database"
	"gin-restapi/internal/user/controllers"
	"gin-restapi/internal/user/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func RouterAuth(r *gin.RouterGroup) {
	db, err := database.ConnectToMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	store := storage.SQLStore(db)

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register(store))
		auth.POST("/login", controllers.Login(store))
	}

}
