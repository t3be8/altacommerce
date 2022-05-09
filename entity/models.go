package entity

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Dob       *time.Time
	Gender    string
	Email     string
	Phone     *string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Addresses []Address
}

func (User) TableName() string {
	return "users"
}

type Address struct {
	ID        uint
	Address   string
	KodePos   int
	UserID    uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
