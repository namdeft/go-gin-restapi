package controllers

import (
	"gin-restapi/internal/common"
	"gin-restapi/internal/token"
	"gin-restapi/internal/user/dto"
	"gin-restapi/internal/user/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func ValidateLoginInput(input *dto.LoginInput) error {
	validate := validator.New()
	err := validate.Struct(input)

	return err
}

func (controller *authController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.LoginInput

		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vErr := ValidateLoginInput(&input)
		if vErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": vErr.Error(),
			})

			return
		}

		tokenKey, err := controller.authService.Login(c.Request.Context(), &input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token.ExtractToken(c)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(tokenKey))

	}
}

func (controller *authController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dto.RegisterInput

		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		vErr := input.ValidateRegisterInput()
		if vErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": vErr.Error(),
			})

			return
		}

		if err := controller.authService.Register(c.Request.Context(), &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(input))
	}
}
