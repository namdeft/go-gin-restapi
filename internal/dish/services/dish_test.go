package services

import (
	"context"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/dto"
	"gin-restapi/internal/dish/services/mocks"
	"gin-restapi/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewDish(t *testing.T) {
	mockRepo := new(mocks.MockRepository)

	dish := dto.DishCreation{
		Name:  "banh da tron",
		Price: "39.99",
	}

	mockRepo.On("CreateDish").Return(&dish, nil)

	testService := NewDishService(mockRepo)
	err := testService.CreateNewDish(context.Background(), &dish)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestGetDishes(t *testing.T) {
	mockRepo := new(mocks.MockRepository)

	dish := model.Dish{ID: 1, Name: "banh da tron", Price: "29.99"}

	mockRepo.On("GetDishes").Return([]model.Dish{dish}, nil)

	testService := NewDishService(mockRepo)

	testPaging := common.Paging{
		Page:  1,
		Limit: 1,
	}

	result, err := testService.GetDishes(context.Background(), &testPaging)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "banh da tron", result[0].Name)
	assert.Equal(t, "29.99", result[0].Price)
	assert.Nil(t, err)
}

func TestGetDish(t *testing.T) {
	mockRepo := new(mocks.MockRepository)

	dish := model.Dish{ID: 1, Name: "banh da tron", Price: "29.99"}

	mockRepo.On("GetDish").Return(&dish, nil)

	testService := NewDishService(mockRepo)

	result, err := testService.GetDish(context.Background(), 1)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "banh da tron", result.Name)
	assert.Equal(t, "29.99", result.Price)
	assert.Nil(t, err)
}

func TestUpdateDish(t *testing.T) {
	mockRepo := new(mocks.MockRepository)

	dish := dto.DishUpdation{
		Name:   "bun dau mam tom",
		Price:  "34.00",
		Status: "available",
	}

	mockRepo.On("UpdateDish").Return(&dish, nil)

	testService := NewDishService(mockRepo)
	err := testService.UpdateDish(context.Background(), 1, &dish)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestDeleteDish(t *testing.T) {
	mockRepo := new(mocks.MockRepository)

	mockRepo.On("DeleteDish").Return(nil, nil)

	testService := NewDishService(mockRepo)

	err := testService.DeleteDish(context.Background(), 1)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}
