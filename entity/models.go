package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar;not null"`
	Dob       *time.Time
	Gender    string
	Email     string  `gorm:"type:varchar(50);not null;unique"`
	Phone     *string `gorm:"type:varchar(16);unique"`
	Password  string
	Products  []Product `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Addresses []Address `gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Address struct {
	gorm.Model
	Address string `gorm:"type:varchar;not null"`
	KodePos int    `gorm:"type:varchar(6);not null"`
	UserID  uint
	Orders  []Order `gorm:"foreignKey:AddressID"`
}

type Product struct {
	gorm.Model
	Name         string
	Description  string
	Price        int
	Status       string
	Stok         int
	Images       string
	CategoryID   uint
	UserID       uint
	OrderDetails []OrderDetail `gorm:"many2many:order_details"`
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

type Order struct {
	gorm.Model
	CartId       uint
	AddressId    uint
	TotalPrice   float64
	TotalQty     int
	OrderDetails []OrderDetail `gorm:"many2many:order_details"`
}

type OrderDetail struct {
	OrderId   uint
	ProductId uint
	Price     float64
	Qty       int
}
