package product

import "github.com/t3be8/altacommerce/entity"

type IProduct interface {
	InsertProduct(newProduct entity.Product) (entity.Product, error)
	SelectProduct() ([]entity.Product, error)
	UpdateProduct(ID int, Stock int) (entity.Product, error)
	DeletedProduct(ID int) (entity.Product, error)
}
