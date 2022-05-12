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
	Addresses []Address `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Address struct {
	gorm.Model
	Address string
	KodePos int
	UserID  uint
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       int
	Status      string
	Stok        int
	Images      string
	CategoryID  uint
	UserID      uint
}

type Category struct {
	gorm.Model
	Name     string
	Products []Product `gorm:"foreignKey:CategoryID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Total     string
	Qty       string
	Users     []User    `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Products  []Product `gorm:"foreignKey:ProductID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
