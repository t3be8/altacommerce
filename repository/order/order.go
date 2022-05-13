package order

import (
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/entity"
	"gorm.io/gorm"
)

type OrderRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		Db: db,
	}
}

func (or *OrderRepo) CreateOrder(order entity.Order) (entity.Order, error) {
	var items []entity.Cart

	if err := or.Db.Where("user_id=?", order.UserID).Find(&items).Error; err != nil {
		log.Warn()
		return order, err
	}
	log.Info(items)

	if len(items) == 0 {
		return order, errors.New("masukkan produk terlebih dahulu")
	}

	var totalPrice float64
	var totalQty int

	for _, val := range items {
		totalPrice += float64(val.Qty) * val.Price
		totalQty += val.Qty
	}

	// products := []entity.Product{}
	// var product entity.Product
	// if err := or.Db.Where("product_id = ?", items[0].ProductID).Find(&product).Error; err != nil {
	// 	return order, err
	// }
	// products = append(products, product)

	var shpCost entity.Shipment
	if err := or.Db.Where("id = ?", order.ShipmentID).Find(&shpCost).Error; err != nil {
		return order, err
	}
	log.Info(shpCost)

	// order.Products = products
	order.TotalQty = totalQty
	order.TotalPrice = totalPrice
	order.TotalPay = totalPrice + shpCost.Cost
	fmt.Println(order)

	if err := or.Db.Create(&order).Error; err != nil {
		return order, err
	}

	// if err := or.Db.Table("products").Where("", order.ShipmentID).Find(&shpCost).Error; err != nil {
	// 	return order, err
	// }

	return order, nil
}

func (or *OrderRepo) CancelOrder(order_id uint) error {
	var order entity.Order

	if err := or.Db.Where("order_id=?", order_id).First(&order).Update("status", "canceled").Error; err != nil {
		log.Warn("Cancel order gagal")
		return err
	}

	return nil
}

func (or *OrderRepo) PayOrder(user_id uint, order_id string) (entity.Order, error) {
	var order entity.Order

	if err := or.Db.Where("user_id = ? AND order_id = ?", user_id, order_id).First(&order).Update("status", "success").Error; err != nil {
		log.Warn("Pembayaran Order Gagal", err)
		return order, errors.New("gagal akses data")
	}

	return order, nil
}
