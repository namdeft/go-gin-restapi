package controllers

import (
	"gin-restapi/internal/favourite/controllers/mocks"
	"gin-restapi/internal/middlewares"
	"gin-restapi/internal/model"
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

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())
	router.POST("/favourites/:dish_id", testController.AddFavourite())

	req := httptest.NewRequest("POST", "/favourites/2", nil)
	assert.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+validToken)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	c.Set("user_id", 1)

	mockService.On("CheckDishExists", 2).Return(nil)
	mockService.On("AddFavourite", mock.Anything, 1, 2).Return(nil)

	router.ServeHTTP(w, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetFavouritesController(t *testing.T) {
	mockService := new(mocks.MockService)
	testController := NewFavouriteController(mockService)

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())
	router.GET("/favourites", testController.GetFavourites())

	req := httptest.NewRequest("GET", "/favourites", nil)
	assert.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+validToken)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	c.Set("user_id", 1)

	dish := []model.Dish{{ID: 1, Name: "Dish 1", Price: "19.99"}}
	mockService.On("GetFavourites", mock.Anything, 1).Return(dish, nil)

	router.ServeHTTP(w, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteFavouriteController(t *testing.T) {
	mockService := new(mocks.MockService)
	testController := favouriteController{favouriteService: mockService}

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())
	router.DELETE("/favourites/:dish_id", testController.DeleteFavourite())

	req := httptest.NewRequest("DELETE", "/favourites/2", nil)
	assert.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+validToken)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	c.Set("user_id", 1)

	mockService.On("DeleteFavourite", mock.Anything, 1, 2).Return(nil)

	router.ServeHTTP(w, req)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
}
