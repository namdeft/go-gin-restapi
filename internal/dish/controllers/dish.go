package controllers

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/dto"
	"gin-restapi/internal/dish/services"
	"gin-restapi/internal/dish/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDishes(store *storage.SqlStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Process()

		sv := services.DishService(store)

		dishes, err := sv.GetDishes(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusAccepted, common.SuccessResponse(dishes, paging, nil))
	}
}

func GetDish(store *storage.SqlStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		sv := services.DishService(store)

		dish, err := sv.GetDish(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(dish))
	}
}

func CreateDish(store *storage.SqlStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.DishCreation

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		vErr := input.ValidateDishCreation()
		if vErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": vErr.Error(),
			})

			return
		}

		sv := services.DishService(store)

		if err := sv.CreateNewDish(c.Request.Context(), &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(input.Id))
	}
}

func UpdateDish(store *storage.SqlStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.DishUpdation

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		vErr := input.ValidateDishUpdation()
		if vErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": vErr.Error(),
			})

			return
		}

		sv := services.DishService(store)

		if err := sv.UpdateDish(c.Request.Context(), id, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(true))
	}
}

func DeleteDish(store *storage.SqlStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		sv := services.DishService(store)

		if err := sv.DeleteDish(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(true))
	}
}
