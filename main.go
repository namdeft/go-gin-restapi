package main

import (
	"gin-restapi/internal/dish"
	"gin-restapi/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	dish.RouterDish(v1)
	user.RouterAuth(v1)

	r.Run()
}
