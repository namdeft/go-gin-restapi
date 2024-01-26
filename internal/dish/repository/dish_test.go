package repository

import (
	"context"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/dto"
	"gin-restapi/internal/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo DishRepository
}

func (s *Suite) SetupSuite() {
	db, mock, err := sqlmock.New()
	require.NoError(s.T(), err)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.db = gormDB
	s.mock = mock
	s.repo = NewDishRepository(gormDB)
}

func (suite *Suite) TestGetDish(t *testing.T) {
	suite.mock.ExpectQuery(`SELECT (.+) FROM dishes WHERE id = ?`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "status"}).AddRow(1, "banh da trong", "29.99", "available"))

	result, err := suite.repo.GetDish(context.Background(), 1)
	assert.Nil(t, err)

	expectedResult := &model.Dish{
		ID:     1,
		Name:   "Test Dish",
		Price:  "19.99",
		Status: model.DishStatusAvailable,
	}
	assert.Equal(t, expectedResult, result)

	assert.Nil(t, suite.mock.ExpectationsWereMet())
}

func (suite *Suite) TestGetDishes(t *testing.T) {
	suite.mock.ExpectQuery(`SELECT (.+) FROM dish`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "status"}).
			AddRow(1, "banh da tron", "29.99", "available").
			AddRow(2, "bun cha", "35.99", "unavailable"))

	suite.mock.ExpectQuery(`SELECT COUNT\(.+\) FROM dish`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	paging := &common.Paging{
		Page:  1,
		Limit: 10,
	}

	result, err := suite.repo.GetDishes(context.Background(), paging)
	assert.Nil(t, err)

	expectedResult := []model.Dish{
		{ID: 1, Name: "banh da tron", Price: "29.99", Status: model.DishStatusAvailable},
		{ID: 2, Name: "bun cha", Price: "35.99", Status: model.DishStatusUnavailable},
	}
	assert.Equal(t, expectedResult, result)

	assert.Nil(t, suite.mock.ExpectationsWereMet())
}

func (suite *Suite) TestDeleteDish(t *testing.T) {
	suite.mock.ExpectExec(`UPDATE dishes SET status = ?, deleted_at = ? WHERE id = ?`).
		WithArgs("deleted", sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.repo.DeleteDish(context.Background(), 1)
	assert.Nil(t, err)

	assert.Nil(t, suite.mock.ExpectationsWereMet())
}

func (suite *Suite) TestCreateDish(t *testing.T) {
	suite.mock.ExpectExec(`INSERT INTO dishes (.+) VALUES (.+)`).
		WithArgs(sqlmock.AnyArg(), "banh da tron", "29.99", "available", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	data := &dto.DishCreation{
		Name:  "banh da tron",
		Price: "29.99",
	}

	err := suite.repo.CreateDish(context.Background(), data)
	assert.Nil(t, err)

	assert.Nil(t, suite.mock.ExpectationsWereMet())
}

func (suite *Suite) TestUpdateDish() {
	suite.mock.ExpectExec(`UPDATE dishes SET (.+) WHERE id = ?`).
		WithArgs(sqlmock.AnyArg(), "Updated Dish", "29.99", "unavailable", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	data := &dto.DishUpdation{
		Name:   "banh da tron",
		Price:  "29.99",
		Status: "unavailable",
	}

	err := suite.repo.UpdateDish(context.Background(), 1, data)
	assert.Nil(suite.T(), err)

	assert.Nil(suite.T(), suite.mock.ExpectationsWereMet())
}

func TestDishRepoSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
