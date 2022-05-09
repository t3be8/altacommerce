package user

import (
	"context"

	"github.com/t3be8/altacommerce/entity"
)

type IUser interface {
	Register(newUser entity.User) (entity.User, error)
	GetAll() ([]entity.User, error)
	IsLogin(ctx context.Context, email, password string) (entity.User, bool, error)
}
