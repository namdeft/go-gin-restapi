package controllers

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/dto"
	"gin-restapi/internal/dish/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DishController interface {
	GetDishes() gin.HandlerFunc
	GetDish() gin.HandlerFunc
	CreateDish() gin.HandlerFunc
	UpdateDish() gin.HandlerFunc
	DeleteDish() gin.HandlerFunc
}

type dishController struct {
	dishService services.DishService
}

func NewDishController(dishService services.DishService) DishController {
	return &dishController{
		dishService: dishService,
	}
}

func (controller *dishController) GetDishes() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Process()

		dishes, err := controller.dishService.GetDishes(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SuccessResponse(dishes, paging, nil))
	}
}

func (controller *dishController) GetDish() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatalf(err.Error())
		}

		dish, err := controller.dishService.GetDish(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dish))
	}
}

func (controller *dishController) CreateDish() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.DishCreation

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		validate := validator.New()
		vErr := validate.Struct(input)

		if vErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": vErr.Error(),
			})

			return
		}

		if err := controller.dishService.CreateNewDish(c.Request.Context(), &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(input.Id))
	}
}

func (controller *dishController) UpdateDish() gin.HandlerFunc {
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

		validate := validator.New()
		vErr := validate.Struct(input)

		if vErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": vErr.Error(),
			})

			return
		}

		if err := controller.dishService.UpdateDish(c.Request.Context(), id, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func (controller *dishController) DeleteDish() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if err := controller.dishService.DeleteDish(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
