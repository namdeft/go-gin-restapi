package main

import (
	"gin-restapi/internal/database"
	"gin-restapi/internal/dish"
	"gin-restapi/internal/dish/controllers"
	"gin-restapi/internal/dish/repository"
	"gin-restapi/internal/dish/services"
	"gin-restapi/internal/favourite"
	favouriteCon "gin-restapi/internal/favourite/controllers"
	favouriteRepo "gin-restapi/internal/favourite/repository"
	favouriteServ "gin-restapi/internal/favourite/services"
	"gin-restapi/internal/user"
	authCon "gin-restapi/internal/user/controllers"
	userRepo "gin-restapi/internal/user/repository"
	authServ "gin-restapi/internal/user/services"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	db, err = database.ConnectToMySQL()

	dishRepository repository.DishRepository  = repository.NewDishRepository(db)
	dishService    services.DishService       = services.NewDishService(dishRepository)
	dishController controllers.DishController = controllers.NewDishController(dishService)

	userRepository userRepo.UserRepository = userRepo.NewUserRepository(db)
	authService    authServ.AuthService    = authServ.NewAuthService(userRepository)
	authController authCon.AuthController  = authCon.NewAuthController(authService)

	favouriteRepository favouriteRepo.FavouriteRepository = favouriteRepo.NewFavouriteRepository(db)
	favouriteService    favouriteServ.FavouriteService    = favouriteServ.NewFavouriteService(favouriteRepository)
	favouriteController favouriteCon.FavouriteController  = favouriteCon.NewFavouriteController(favouriteService)
)

func main() {
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")
	dish.RouterDish(v1, dishController)
	user.RouterAuth(v1, authController)
	favourite.RouterFavourite(v1, favouriteController)
	r.Run()
}
