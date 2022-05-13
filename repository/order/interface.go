package order

import "github.com/t3be8/altacommerce/entity"

type IOrder interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	CancelOrder(order_id int) error
	PayOrder(user_id uint, order_id string) (entity.Order, error)
}
