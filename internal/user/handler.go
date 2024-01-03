package user

import (
	"gin-restapi/internal/database"
	"gin-restapi/internal/user/transport"
	"log"

	"github.com/gin-gonic/gin"
)

func RouterAuth(r *gin.RouterGroup) {
	db, err := database.ConnectToMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", transport.Register(db))
		auth.POST("/login", transport.Login(db))
	}

}
