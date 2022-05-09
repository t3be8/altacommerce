package product

import (
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/entity"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *ProductRepo {
	return &ProductRepo{
		Db: db,
	}
}

type ProductRepo struct {
	Db *gorm.DB
}

func (pr *ProductRepo) InsertProduct(newProduct entity.Product) (entity.Product, error) {
	if err := pr.Db.Create(&newProduct).Error; err != nil {
		log.Warn(err)
		return entity.Product{}, errors.New("tidak bisa insert")
	}
	log.Info()
	return newProduct, nil
}

func (pr *ProductRepo) SelectProduct() ([]entity.Product, error) {
	arrProduct := []entity.Product{}

	if err := pr.Db.Find(&arrProduct).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa select")
	}

	if len(arrProduct) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}

	log.Info()
	return arrProduct, nil
}

func (pr *ProductRepo) UpdateProduct(ID int, Stock int) (entity.Product, error) {
	product := entity.Product{}
	if err := pr.Db.Model(&product).Where("id = ?", ID).Update("stok", Stock).Error; err != nil {
		log.Warn(err)
		return entity.Product{}, errors.New("tidak bisa update")
	}
	log.Info()
	return product, nil
}

func (pr *ProductRepo) DeletedProduct(ID int) (entity.Product, error) {
	product := entity.Product{}
	if err := pr.Db.Model(&product).Where("id = ?", ID).Delete(&product).Error; err != nil {
		log.Warn(err)
		return entity.Product{}, errors.New("tidak bisa delete")
	}
	log.Info()
	return product, nil
}
