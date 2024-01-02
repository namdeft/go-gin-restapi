package main

import (
	"errors"
	"fmt"
	"gin-restapi/internal/common"
	"gin-restapi/internal/dish/model"
	"gin-restapi/internal/ingredient/model"
	"gin-restapi/internal/middlewares"
	"gin-restapi/internal/token"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/html"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// type DishStatus int

// const (
// 	DishStatusUnavailable DishStatus = iota
// 	DishStatusAvailable
// 	DishStatusDeleted
// )

// type Dish struct {
// 	ID         int        `json:"id"`
// 	Name       string     `json:"name"`
// 	Price      string     `json:"price"`
// 	Status     DishStatus `json:"status"`
// 	Updated_At time.Time  `json:"updated_at"`
// 	Created_At time.Time  `json:"created_at"`
// 	Deleted_At time.Time  `json:"deleted_at"`
// }

// func (Dish) TableName() string {
// 	return "dish"
// }

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

// type Ingredient struct {
// 	ID            string    `json:"id"`
// 	Name          string    `json:"name"`
// 	Quantity      int       `json:"quantity"`
// 	Import_Date   time.Time `json:"import_date"`
// 	Export_Date   time.Time `json:"export_date"`
// 	Counting_Unit int       `json:"counting_unit"`
// }

var allDishStatuses = [3]string{
	"unavailable",
	"available",
	"deleted",
}

func (dish DishStatus) String() string {
	return allDishStatuses[dish]
}

func parseStrToDishStatus(s string) model.DishStatus {
	for i := range allDishStatuses {
		if allDishStatuses[i] == s {
			return model.DishStatus(i)
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

func getIngredients(c *gin.Context) {
	var ingredients []model.Ingredient
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

type LoginInput struct {
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

func GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {
	var user User

	if err := DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return "", nil
	}

	err := VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", nil
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{
		Email:    input.Email,
		Password: input.Password,
	}

	tokenKey, err := LoginCheck(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token.ExtractToken(c)

	c.JSON(http.StatusBadRequest, common.SimpleSuccessResponse(tokenKey))
}

func main() {
	connectToMySQL()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		dishes := v1.Group("/dishes")
		dishes.Use(middlewares.JwtAuthMiddleware())
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
			auth.POST("/login", login)
		}
	}

	r.Run()
}
