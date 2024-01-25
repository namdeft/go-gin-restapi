package controllers

import (
	"gin-restapi/internal/favourite/controllers/mocks"
	"gin-restapi/internal/favourite/model"
	"gin-restapi/internal/middlewares"
	"gin-restapi/internal/token"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddFavourite(t *testing.T) {
	mockService := new(mocks.MockService)
	testController := NewFavouriteController(mockService)

	mockService.On("AddFavourite", mock.Anything, 1, 2).Return(nil, nil)

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())

	req := httptest.NewRequest("POST", "/favourites/1/2", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	router.Use(func(c *gin.Context) {
		c.Set("user_id", 1)
		c.Set("dish_id", 2)
	})
	router.POST("/favourites/:user_id/:dish_id", testController.AddFavourite())
	router.ServeHTTP(w, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetFavouritesController(t *testing.T) {
	mockService := new(mocks.MockService)
	testController := favouriteController{favouriteService: mockService}

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	dish := []model.Dish{{ID: 1, Name: "Dish 1", Price: "19.99"}}
	mockService.On("GetFavourites", mock.Anything, 1).Return(dish, nil)

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())

	req := httptest.NewRequest("GET", "/favourites/1", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	router.Use(func(c *gin.Context) {
		c.Set("user_id", 1)
	})
	router.GET("/favourites/:user_id", testController.GetFavourites())
	router.ServeHTTP(w, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteFavouriteController(t *testing.T) {
	mockService := new(mocks.MockService)
	testController := favouriteController{favouriteService: mockService}

	mockService.On("DeleteFavourite", mock.Anything, 1, 2).Return(nil, nil)

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())

	req := httptest.NewRequest("DELETE", "/favourites/1/2", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	router.Use(func(c *gin.Context) {
		c.Set("user_id", 1)
		c.Set("dish_id", 2)
	})
	router.DELETE("/favourites/:user_id/:dish_id", testController.DeleteFavourite())
	router.ServeHTTP(w, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}
