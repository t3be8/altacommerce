package cart

import (
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/entity"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *CartRepo {
	return &CartRepo{
		Db: db,
	}
}

type CartRepo struct {
	Db *gorm.DB
}

func (cr *CartRepo) InsertCart(newCart entity.Cart) (entity.Cart, error) {
	if err := cr.Db.Create(&newCart).Error; err != nil {
		log.Warn(err)
		return entity.Cart{}, errors.New("tidak bisa insert")
	}
	log.Info()
	return newCart, nil
}

func (cr *CartRepo) SelectCart() ([]entity.Cart, error) {
	arrCart := []entity.Cart{}

	if err := cr.Db.Find(&arrCart).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select")
	}

	if len(arrCart) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrCart, nil
}

func (cr *CartRepo) UpdateCart(ID int, Total int) (entity.Cart, error) {
	cart := entity.Cart{}
	if err := cr.Db.Model(&cart).Where("id = ?", ID).Update("stok", Total).Error; err != nil {
		log.Warn(err)
		return entity.Cart{}, errors.New("tidak bisa update")
	}
	log.Info()
	return cart, nil
}

func (cr *CartRepo) DeletedCart(ID int) (entity.Cart, error) {
	cart := entity.Cart{}
	if err := cr.Db.Model(&cart).Where("id = ?", ID).Delete(&cart).Error; err != nil {
		log.Warn(err)
		return entity.Cart{}, errors.New("tidak bisa delete")
	}
	log.Info()
	return cart, nil
}
