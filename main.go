package main

import (
	"errors"
	"fmt"
	"gin-restapi/common"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/html"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DishStatus int

const (
	DishStatusUnavailable DishStatus = iota
	DishStatusAvailable
	DishStatusDeleted
)

type Dish struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Price      string     `json:"price"`
	Status     DishStatus `json:"status"`
	Updated_At time.Time  `json:"updated_at"`
	Created_At time.Time  `json:"created_at"`
	Deleted_At time.Time  `json:"deleted_at"`
}

func (Dish) TableName() string {
	return "dish"
}

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

type Ingredient struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Quantity      int       `json:"quantity"`
	Import_Date   time.Time `json:"import_date"`
	Export_Date   time.Time `json:"export_date"`
	Counting_Unit int       `json:"counting_unit"`
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
			"error": err,
		})

		return
	}

	paging.Process()

	if err := DB.Table(Dish{}.TableName()).Count(&paging.Total); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var dishes []Dish
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
	var dish Dish
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
	var dish Dish

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

	if err := DB.Table(Dish{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
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

func getIngredients(c *gin.Context) {
	var ingredients []Ingredient
	if err := DB.Find(&ingredients).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(ingredients))
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) BeforeSaveUser() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	return nil
}

func (u *User) SaveUser() error {
	if err := DB.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := u.BeforeSaveUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.SaveUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(input))
}

func main() {
	connectToMySQL()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		dishes := v1.Group("/dishes")
		{
			dishes.GET("/", getDishes)
			dishes.GET("/:id", getDish)
			dishes.POST("/", createDish)
			dishes.PUT("/:id", updateDish)
			dishes.DELETE("/:id", deleteDishSoftly)
		}
		ingredients := v1.Group("/ingredients")
		{
			ingredients.GET("/", getIngredients)
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/register", register)
		}
	}

	r.Run()
}
