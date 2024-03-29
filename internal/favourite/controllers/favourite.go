package controllers

import (
	"errors"
	"gin-restapi/internal/common"
	"gin-restapi/internal/favourite/services"
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

func NewFavouriteController(favouriteService services.FavouriteService) FavouriteController {
	return &favouriteController{
		favouriteService: favouriteService,
	}
}

func (controller favouriteController) AddFavourite() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": errors.New("User does not exist"),
			})
		}

		dishId, err := strconv.Atoi(c.Param("dish_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if err := controller.favouriteService.CheckDishExists(dishId); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := controller.favouriteService.AddFavourite(c.Request.Context(), userId.(int), dishId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
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
		userId, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": error.Error,
			})
		}

		dishes, err := controller.favouriteService.GetFavourites(c.Request.Context(), userId.(int))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dishes))
	}
}

func (controller favouriteController) DeleteFavourite() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": errors.New("User does not exist"),
			})
		}

		dishId, err := strconv.Atoi(c.Param("dish_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if err := controller.favouriteService.DeleteFavourite(c.Request.Context(), userId.(int), dishId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dishId))
	}
}
