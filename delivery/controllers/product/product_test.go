package product

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/t3be8/altacommerce/entity"
)

var token string

func TestSelectProduct(t *testing.T) {
	t.Run("berhasil select product", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product")

		productController := New(&mockProductRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("RH$SI4")})(productController.SelectProduct())(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entity.Product
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, resp.Data[0].Name, "Asus ROG")
		// assert.Equal(t, 500, resp.Code)
	})
	// t.Run("Error select product", func(t *testing.T) {
	// 	e := echo.New()
	// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	res := httptest.NewRecorder()
	// 	context := e.NewContext(req, res)
	// 	context.SetPath("/product")

	// 	productController := New(&errorMockProductRepository{}, validator.New())
	// 	productController.SelectProduct()

	// 	type response struct {
	// 		Code    int
	// 		Message string
	// 		Status  bool
	// 		Data    []entity.Product
	// 	}

	// 	var resp response

	// 	json.Unmarshal([]byte(res.Body.Bytes()), &resp)
	// 	assert.Nil(t, resp.Data)
	// 	assert.False(t, resp.Status)
	// 	assert.Equal(t, 500, resp.Code)
	// })
}

// dummy data

type mockProductRepository struct{}

func (mpr *mockProductRepository) SelectProduct() ([]entity.Product, error) {
	return []entity.Product{{Name: "Asus ROG", Description: "ini laptop gaming", Price: 1200000, Stok: 100, Images: "image.com"}}, nil
}

func (mpr *mockProductRepository) InsertProduct(newProduct entity.Product) (entity.Product, error) {
	return newProduct, nil
}

func (mpr *mockProductRepository) UpdateProduct(ID int, Stock int) (entity.Product, error) {
	return entity.Product{}, nil
}

func (mpr *mockProductRepository) DeletedProduct(ID int) (entity.Product, error) {
	return entity.Product{}, nil
}

func (mpr *mockProductRepository) GetAllById(ID int) ([]entity.Product, error) {
	return []entity.Product{{Name: "Asus ROG", Description: "ini laptop gaming", Price: 1200000, Stok: 100, Images: "image.com"}}, nil
}

func (mpr *mockProductRepository) GetAllByCategory(ID int) ([]entity.Product, error) {
	return []entity.Product{{Name: "Asus ROG", Description: "ini laptop gaming", Price: 1200000, Stok: 100, Images: "image.com"}}, nil
}

// type errorMockProductRepository struct{}

// func (empr *errorMockProductRepository) SelectProduct() ([]entity.Product, error) {
// 	return nil, errors.New(("tidak bisa select product"))
// }

// func (empr *errorMockProductRepository) InsertProduct(newProduct entity.Product) (entity.Product, error) {
// 	return entity.Product{}, errors.New(("tidak bisa Insert product"))
// }

// func (empr *errorMockProductRepository) UpdateProduct(ID int, Stock int) (entity.Product, error) {
// 	return entity.Product{}, errors.New(("tidak bisa select product"))
// }

// func (empr *errorMockProductRepository) DeletedProduct(ID int) (entity.Product, error) {
// 	return entity.Product{}, errors.New(("tidak bisa select product"))
// }

// func (empr *errorMockProductRepository) GetAllById(ID int) ([]entity.Product, error) {
// 	return []entity.Product{}, errors.New(("tidak bisa select product"))
// }

// func (empr *errorMockProductRepository) GetAllByCategory(ID int) ([]entity.Product, error) {
// 	return []entity.Product{}, errors.New(("tidak bisa select product"))
// }
