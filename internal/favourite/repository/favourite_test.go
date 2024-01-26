package repository

import (
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
	repo FavouriteRepository
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
	s.repo = NewFavouriteRepository(gormDB)
}

func (suite *Suite) TestAddFavourite() {
	userId := 1
	dishId := 2

	suite.mock.ExpectQuery(`SELECT (.+) FROM "users" WHERE "id" = (.+)`).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(userId, "namdeft20@gmail.com"))

	suite.mock.ExpectQuery(`SELECT (.+) FROM "dishes" WHERE "id" = (.+)`).
		WithArgs(dishId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "status"}).AddRow(dishId, "banh da tron", "29.99", "available"))

	suite.mock.ExpectExec(`INSERT INTO "favourite" (.+) VALUES (.+)`).
		WithArgs(userId, dishId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.repo.AddFavourite(userId, dishId)
	assert.Nil(suite.T(), err)
}

func (suite *Suite) TestGetFavourites() {
	userId := 1

	suite.mock.ExpectQuery(`SELECT (.+) FROM "users" WHERE "id" = (.+)`).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(userId, "namdeft20@gmail.com"))

	suite.mock.ExpectQuery(`SELECT (.+) FROM "dishes" INNER JOIN "favourite" ON "dishes"."id" = "favourite"."dish_id" WHERE "favourite"."user_id" = (.+)`).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "status"}).AddRow(1, "banh da tron", "29.99", "available"))

	dishes, err := suite.repo.GetFavourites(userId)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), dishes, 1)
	assert.Equal(suite.T(), "banh da tron", dishes[0].Name)
}

func (suite *Suite) TestDeleteFavourite() {
	userId := 1
	dishId := 2

	suite.mock.ExpectQuery(`SELECT (.+) FROM "users" WHERE "id" = (.+)`).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(userId, "namdeft20@gmail.com"))

	suite.mock.ExpectQuery(`SELECT (.+) FROM "dishes" INNER JOIN "favourite" ON "dishes"."id" = "favourite"."dish_id" WHERE "favourite"."user_id" = (.+) AND "favourite"."dish_id" = (.+)`).
		WithArgs(userId, dishId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "status"}).AddRow(dishId, "banh da tron", "29.99", "available"))

	suite.mock.ExpectExec(`DELETE FROM "favourite" WHERE "user_id" = (.+) AND "dish_id" = (.+)`).
		WithArgs(userId, dishId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.repo.DeleteFavourite(userId, dishId)
	assert.Nil(suite.T(), err)
}
func (suite *Suite) TestCheckDishExists() {
	dishId := 1

	suite.mock.ExpectQuery(`SELECT (.+) FROM "dishes" WHERE "id" = (.+)`).
		WithArgs(dishId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "status"}).AddRow(dishId, "banh da tron", "29.99", "available"))

	err := suite.repo.CheckDishExists(dishId)
	assert.Nil(suite.T(), err)
}

func TestFavouriteRepoSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
