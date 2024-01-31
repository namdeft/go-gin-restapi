package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Dishes   []Dish `gorm:"many2many:favourite" json:"-"`
}
