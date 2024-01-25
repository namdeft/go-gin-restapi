package services

import (
	"context"
	"gin-restapi/internal/favourite/model"
	"gin-restapi/internal/favourite/services/mocks"
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
