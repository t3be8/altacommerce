package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255);not null"`
	Dob       *time.Time
	Gender    string
	Email     string  `gorm:"type:varchar(50);not null;unique"`
	Phone     *string `gorm:"type:varchar(16);unique"`
	Password  string
	Products  []Product `gorm:"foreignKey:UserID"`
	Addresses []Address `gorm:"foreignKey:UserID"`
	Carts     []Cart    `gorm:"foreignKey:UserID"`
	Orders    []Order   `gorm:"foreignKey:UserID"`
}

type Address struct {
	gorm.Model
	Address string `gorm:"type:text;not null"`
	KodePos int    `gorm:"type:varchar(6);not null"`
	UserID  uint
	// Orders  []Order `gorm:"foreignKey:AddressID"`
}

type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)"`
	Description string
	Price       int
	Stok        int
	Images      string
	CategoryID  uint
	UserID      uint
}

type Category struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(25);not null"`
	Products []Product `gorm:"foreignKey:CategoryID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Price     float64 `gorm:"type:decimal(18,2);not null"`
	Qty       int
}

type Order struct {
	gorm.Model
	UserID     uint
	CartID     uint
	Address    string
	ShipmentID uint
	Status     string `gorm:"type:varchar(10);default:'waiting'"`
	TotalQty   int
	TotalPrice float64   `gorm:"type:decimal(18,2);not null"`
	TotalPay   float64   `gorm:"type:decimal(18,2);not null"`
	Products   []Product `gorm:"many2many:order_details;"`
}

type Shipment struct {
	gorm.Model
	Type   string  `gorm:"type:varchar(10);not null"`
	Cost   float64 `gorm:"type:numeric(18,2);not null"`
	Orders []Order `gorm:"foreignKey:ShipmentID"`
}

type OrderDetail struct {
	OrderID   uint
	ProductID uint
	Price     float64
	Qty       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
