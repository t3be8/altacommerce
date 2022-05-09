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
