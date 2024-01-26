package controllers

import (
	"bytes"
	"gin-restapi/internal/dish/controllers/mocks"
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

func TestGetDishes(t *testing.T) {
	mockService := new(mocks.MockService)

	dish := []model.Dish{{ID: 1, Name: "banh da tron", Price: "29.99"}}
	mockService.On("GetDishes", mock.Anything, mock.Anything).Return(dish, nil)

	testController := NewDishController(mockService)

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())
	router.GET("/dishes", testController.GetDish())
	ts := httptest.NewServer(router)

	req := httptest.NewRequest(http.MethodGet, ts.URL+"/dishes", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)
	q := req.URL.Query()
	q.Add("first", "limit")
	q.Add("second", "page")

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestGetDish(t *testing.T) {
	mockService := new(mocks.MockService)

	dish := &model.Dish{ID: 1, Name: "banh da tron", Price: "29.99"}
	mockService.On("GetDish", mock.Anything, 1).Return(dish, nil)

	testController := NewDishController(mockService)

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())
	router.GET("/dishes/:id", testController.GetDish())
	ts := httptest.NewServer(router)

	req := httptest.NewRequest(http.MethodGet, ts.URL+"/dishes/1", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestCreateDish(t *testing.T) {
	mockService := new(mocks.MockService)

	mockService.On("CreateNewDish", mock.Anything, mock.Anything).Return(nil, nil)

	testController := NewDishController(mockService)

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())
	router.POST("/dishes", testController.CreateDish())
	ts := httptest.NewServer(router)

	payload := []byte(`{"name": "banh da tron", "price": "29.99"}`)

	req := httptest.NewRequest(http.MethodPost, ts.URL+"/dishes", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", "Bearer "+validToken)

	res, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockService.AssertExpectations(t)
}

func TestUpdateDish(t *testing.T) {
	mockService := new(mocks.MockService)

	mockService.On("UpdateDish", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	testController := NewDishController(mockService)

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())
	router.PUT("/dishes/:id", testController.UpdateDish())
	ts := httptest.NewServer(router)

	payload := []byte(`{"name": "banh da tron", "price": "29.99"}`)

	req := httptest.NewRequest(http.MethodPut, ts.URL+"/dishes/1", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", "Bearer "+validToken)

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestDeleteDish(t *testing.T) {
	mockService := new(mocks.MockService)

	mockService.On("DeleteDish", mock.Anything, mock.Anything).Return(nil, nil)

	testController := NewDishController(mockService)

	validToken, err := token.GenerateToken(1)
	if err != nil {
		t.Fatal("Error generating token:", err)
	}

	router := gin.Default()
	router.Use(middlewares.JwtAuthMiddleware())
	router.DELETE("/dishes/:id", testController.DeleteDish())
	ts := httptest.NewServer(router)

	req := httptest.NewRequest(http.MethodDelete, ts.URL+"/dishes/1", nil)
	req.Header.Set("Authorization", "Bearer "+validToken)

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}
