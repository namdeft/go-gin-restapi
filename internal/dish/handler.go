package dish

import (
	"errors"
	"fmt"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DishStatus int

const (
	DishStatusUnavailable DishStatus = iota
	DishStatusAvailable
	DishStatusDeleted
)

type DishCreation struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func (DishCreation) TableName() string {
	return "dish"
}

type DishUpdation struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	Status string `json:"status"`
}

func (DishUpdation) TableName() string {
	return "dish"
}

var allDishStatuses = [3]string{
	"unavailable",
	"available",
	"deleted",
}

func (dish DishStatus) String() string {
	return allDishStatuses[dish]
}

func parseStrToDishStatus(s string) DishStatus {
	for i := range allDishStatuses {
		if allDishStatuses[i] == s {
			return DishStatus(i)
		}
	}

	return DishStatusUnavailable
}

func (dish *DishStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql1: %s", value))
	}

	v := parseStrToDishStatus(string(bytes))

	*dish = v

	return nil
}

func (dish *DishStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", dish.String())), nil
}

var DB *gorm.DB

func connectToMySQL() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Some error with env variable: ", err)
	}

	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
}

func getDishes(c *gin.Context) {
	var paging common.Paging

	if err := c.ShouldBind(&paging); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	paging.Process()

	if err := DB.Table(model.Dish{}.TableName()).Count(&paging.Total); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var dishes []model.Dish
	if err := DB.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&dishes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, common.SuccessResponse(dishes, paging, nil))
}

func getDish(c *gin.Context) {
	var dish model.Dish
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := DB.Where("id = ?", id).First(&dish).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(dish))
}

func createDish(c *gin.Context) {
	var input DishCreation

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	if err := DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(input.Id))
}

func updateDish(c *gin.Context) {
	var input DishUpdation

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := DB.Where("id = ?", id).Updates(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(true))
}

func deleteDish(c *gin.Context) {
	var dish model.Dish

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := DB.Where("id = ?", id).Delete(&dish).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(true))
}

func deleteDishSoftly(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := DB.Table(model.Dish{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status":     "deleted",
		"deleted_at": time.Now(),
	}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(true))
}
