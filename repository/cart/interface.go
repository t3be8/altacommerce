package cart

import "github.com/t3be8/altacommerce/entity"

type ICart interface {
	InsertCart(newCart entity.Cart) (entity.Cart, error)
	SelectCart() ([]entity.Cart, error)
	UpdateCart(ID int, Stock int) (entity.Cart, error)
	DeletedCart(ID int) (entity.Cart, error)
}
