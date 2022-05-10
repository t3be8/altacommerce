package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	view "github.com/t3be8/altacommerce/delivery/views"
	"github.com/t3be8/altacommerce/entity"

	cartView "github.com/t3be8/altacommerce/delivery/views/cart"
	cartRepo "github.com/t3be8/altacommerce/repository/cart"
)

type CartController struct {
	Repo  cartRepo.ICart
	Valid *validator.Validate
}

func New(repo cartRepo.ICart, valid *validator.Validate) *CartController {
	return &CartController{
		Repo:  repo,
		Valid: valid,
	}
}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return userId
	}
	return 0
}

func (cc *CartController) InsertCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpCart cartView.InsertCartRequest

		if err := c.Bind(&tmpCart); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusBadRequest, cartView.BadRequest())
		}

		if err := cc.Valid.Struct(tmpCart); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusBadRequest, cartView.BadRequest())
		}

		res, err := cc.Repo.InsertCart(entity.Cart{TotalQty: tmpCart.TotalQty})

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("berhasil insert cart")
		return c.JSON(http.StatusCreated, cartView.SuccessInsert(res))
	}
}

func (cc *CartController) SelectCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := ExtractTokenUserId(c)

		fmt.Println(id)
		res, err := cc.Repo.SelectCart()

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("berhasil select Cart")
		return c.JSON(http.StatusOK, cartView.SuccessSelect(res))
	}
}

func (cc *CartController) UpdateCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpCart cartView.UpdateCartRequest
		id := ExtractTokenUserId(c)

		fmt.Println(id)

		res, err := cc.Repo.UpdateCart(tmpCart.ID, tmpCart.Stok)
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		log.Info("berhasil update cart")
		return c.JSON(http.StatusCreated, cartView.SuccessUpdate(res))
	}
}

func (cc *CartController) DeletedCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpCart cartView.DeleteCartRequest
		id := ExtractTokenUserId(c)

		fmt.Println(id)
		res, err := cc.Repo.DeletedCart(tmpCart.ID)
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		log.Info("berhasi delete product")
		return c.JSON(http.StatusOK, cartView.SuccessDelete(res))
	}
}
