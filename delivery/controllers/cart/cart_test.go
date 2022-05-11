package cart

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

func TestSelectCart(t *testing.T) {
	t.Run("berhasil select cart", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart")

		cartController := New(&mockCartRepository{}, validator.New())
		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("RH$SI4")})(cartController.SelectCart())(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entity.Cart
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, resp.Data[0].Total, 10000)
		// assert.Equal(t, 500, resp.Code)
	})

}
func TestInsertCart(t *testing.T) {

}

func TestUpdateCart(t *testing.T) {

}

func TestDeleteCart(t *testing.T) {

}

// dummy data

type mockCartRepository struct{}

func (mcr *mockCartRepository) SelectCart() ([]entity.Cart, error) {
	return []entity.Cart{{Total: 10000, Qty: 5}}, nil
}

func (mcr *mockCartRepository) InsertCart(newCart entity.Cart) (entity.Cart, error) {
	return newCart, nil
}

func (mpr *mockCartRepository) UpdateCart(ID int, Stock int) (entity.Cart, error) {
	return entity.Cart{}, nil
}

func (mpr *mockCartRepository) DeletedCart(ID int) (entity.Cart, error) {
	return entity.Cart{}, nil
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
