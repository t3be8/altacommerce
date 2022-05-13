package order

import (
	"errors"

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

	if err := or.Db.Where("user_id = ?", order.UserID).Find(&items).Error; err != nil {
		return order, err
	}

	if len(items) == 0 {
		return order, errors.New("masukkan produk terlebih dahulu")
	}

	var totalPrice float64
	var totalQty int

	for _, v := range items {
		totalPrice += float64(v.Qty) * v.Price
		totalQty += v.Qty
	}

	products := []entity.Product{}
	var product entity.Product
	if err := or.Db.Where("product_id = ?", items[0].ProductID).Find(&product).Error; err != nil {
		return order, err
	}
	products = append(products, product)

	var shpCost entity.Shipment
	if err := or.Db.Where("shipment_id = ?", order.ShipmentID).Find(&shpCost).Error; err != nil {
		return order, err
	}

	order.Products = products
	order.TotalQty = totalQty
	order.TotalPrice = totalPrice
	order.TotalPay = totalPrice + shpCost.Cost

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
