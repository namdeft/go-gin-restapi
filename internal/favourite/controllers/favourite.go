package controllers

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/favourite/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavouriteController interface {
	AddFavourite() gin.HandlerFunc
	GetFavourites() gin.HandlerFunc
	DeleteFavourite() gin.HandlerFunc
}

type favouriteController struct {
	favouriteService services.FavouriteService
}

func NewDishController(favouriteService services.FavouriteService) FavouriteController {
	return &favouriteController{
		favouriteService: favouriteService,
	}
}

func (controller favouriteController) AddFavourite() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		dishId, err := strconv.Atoi(c.Param("dish_id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		if err := controller.favouriteService.AddFavourite(c.Request.Context(), userId, dishId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"UserId": userId,
			"DishId": dishId,
		})
	}
}

func (controller favouriteController) GetFavourites() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		dishes, err := controller.favouriteService.GetFavourites(c.Request.Context(), userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dishes))
	}
}

func (controller favouriteController) DeleteFavourite() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		dishId, err := strconv.Atoi(c.Param("dish_id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		if err := controller.favouriteService.DeleteFavourite(c.Request.Context(), userId, dishId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dishId))
	}
}
