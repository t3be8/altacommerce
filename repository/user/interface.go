package user

import (
	"github.com/t3be8/altacommerce/entity"
)

type IUser interface {
	Register(newUser entity.User) (entity.User, error)
	// GetAll() ([]entity.User, error)
	IsLogin(email, password string) (entity.User, bool, error)
}
