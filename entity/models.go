package entity

import (
	"gorm.io/gorm"
)

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
