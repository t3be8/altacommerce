package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255)"`
	Dob       *time.Time
	Gender    string
	Email     string
	Phone     *string
	Password  string
	Addresses []Address
}

type Address struct {
	gorm.Model
	Address string
	KodePos int
	UserID  uint
}

type Product struct {
	gorm.Model
	Name             string
	Description      string
	Price            int
	Status           string
	Stok             int
	Image            string
	ProducCategoryID uint
	UserID           uint
	ProductCategory  []ProductCategory
}

type ProductCategory struct {
	gorm.Model
	Name string
}

type Cart struct {
	gorm.Model
	TotalQty  int
	UserID    uint
	ProductID uint
	User      []User
	Product   []Product
}
