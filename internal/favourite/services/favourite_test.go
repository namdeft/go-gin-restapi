package services

import (
	"context"
	"errors"
	"gin-restapi/internal/favourite/services/mocks"
	"gin-restapi/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddFavourite(t *testing.T) {
	mockRepo := new(mocks.MockRepository)

	mockRepo.On("AddFavourite", mock.Anything, 1, 2).Return(nil, nil)

	testService := NewFavouriteService(mockRepo)

	err := testService.AddFavourite(context.Background(), 1, 2)

	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestGetFavourites(t *testing.T) {
	mockRepo := new(mocks.MockRepository)

	expectedDishes := []model.Dish{{ID: 1, Name: "banh da tron", Price: "29.99"}}
	mockRepo.On("GetFavourites", mock.Anything, 1).Return(expectedDishes, nil)

	testService := NewFavouriteService(mockRepo)

	result, err := testService.GetFavourites(context.Background(), 1)

	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, expectedDishes, result)
}

func TestDeleteFavourite(t *testing.T) {
	mockRepo := new(mocks.MockRepository)

	mockRepo.On("DeleteFavourite", mock.Anything, 1, 2).Return(nil, nil)

	testService := NewFavouriteService(mockRepo)

	err := testService.DeleteFavourite(context.Background(), 1, 2)

	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestCheckDishExists(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := NewFavouriteService(mockRepo)

	mockRepo.On("CheckDishExists", 1).Return(nil)
	err := service.CheckDishExists(1)
	assert.Nil(t, err)

	expectedError := errors.New("Dish not found")
	mockRepo.On("CheckDishExists", 2).Return(expectedError)
	err = service.CheckDishExists(2)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}
