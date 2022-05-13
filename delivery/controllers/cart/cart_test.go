package cart

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/t3be8/altacommerce/entity"
)

var token string

func TestSelectCart(t *testing.T) {
	t.Run("berhasil select cart", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart")

		cartController := New(&mockCartRepository{}, validator.New())
		cartController.SelectCart()(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entity.Cart
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		// log.Fatal(resp)
		assert.Equal(t, 200, resp.Code)
		assert.Equal(t, resp.Data[0].Price, 10000)
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
	return []entity.Cart{{Price: 10000, Qty: 5}}, nil
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
