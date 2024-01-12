package controllers

import (
	"bytes"
	"gin-restapi/internal/dish/controllers/mocks"
	"gin-restapi/internal/dish/model"
	"net/http"
	"net/http/httptest"
	"strings"
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

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/dishes", strings.NewReader(""))
	c.Query("page")
	c.Query("limit")

	testController.GetDishes()(c)

	mockService.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), "banh da tron")
}

func TestGetDish(t *testing.T) {
	mockService := new(mocks.MockService)

	dish := &model.Dish{ID: 1, Name: "Dish1", Price: "10.99"}
	mockService.On("GetDish", mock.Anything, 1).Return(dish, nil)

	testController := NewDishController(mockService)

	// Create a test HTTP server
	router := gin.Default()
	router.GET("/dish/:id", testController.GetDish())
	ts := httptest.NewServer(router)

	// Make a request to the endpoint
	resp, err := http.Get(ts.URL + "/dish/1")
	assert.Nil(t, err)

	// Assert the response status code and body
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	// You may need to parse the JSON response body and compare it with expectedDish

	// Assert that the expectations were met
	mockService.AssertExpectations(t)
}

func TestCreateDish(t *testing.T) {
	mockService := new(mocks.MockService)

	mockService.On("CreateNewDish", mock.Anything, mock.Anything).Return(nil, nil)

	// Create an instance of the DishController with the mock service
	testController := NewDishController(mockService)

	// Create a test HTTP server
	router := gin.Default()
	router.POST("/dishes", testController.CreateDish())
	ts := httptest.NewServer(router)

	// Prepare a JSON payload for the request
	payload := []byte(`{"name": "banh da tron", "price": "29.99"}`)

	// Make a request to the endpoint with the JSON payload
	res, err := http.Post(ts.URL+"/dishes", "application/json", bytes.NewBuffer(payload))
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockService.AssertExpectations(t)
}

func TestUpdateDish(t *testing.T) {
	mockService := new(mocks.MockService)

	mockService.On("UpdateDish", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	testController := NewDishController(mockService)

	router := gin.Default()
	router.PUT("/dishes/:id", testController.UpdateDish())
	ts := httptest.NewServer(router)

	// Prepare a JSON payload for the request
	payload := []byte(`{"name": "banh da tron", "price": "29.99"}`)

	// Make a request to the endpoint with the JSON payload
	req, err := http.NewRequest("PUT", ts.URL+"/dishes/1", bytes.NewBuffer(payload))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestDeleteDish(t *testing.T) {
	mockService := new(mocks.MockService)

	mockService.On("DeleteDish", mock.Anything, mock.Anything).Return(nil, nil)

	// Create an instance of the DishController with the mock service
	testController := NewDishController(mockService)

	// Create a test HTTP server
	router := gin.Default()
	router.DELETE("/dishes/:id", testController.DeleteDish())
	ts := httptest.NewServer(router)

	// Make a request to the endpoint
	req, err := http.NewRequest("DELETE", ts.URL+"/dishes/1", nil)
	assert.Nil(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Assert that the expectations were met
	mockService.AssertExpectations(t)
}
