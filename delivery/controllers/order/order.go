package order

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/t3be8/altacommerce/delivery/middlewares"
	view "github.com/t3be8/altacommerce/delivery/views"
	vo "github.com/t3be8/altacommerce/delivery/views/order"
	"github.com/t3be8/altacommerce/entity"
	"github.com/t3be8/altacommerce/repository/order"
	"github.com/t3be8/altacommerce/utils"
)

type OrderController struct {
	Repo     order.IOrder
	Midtrans utils.ConfigMidtrans
	Valid    *validator.Validate
}

func New(repo order.IOrder, midtrans utils.ConfigMidtrans, valid *validator.Validate) *OrderController {
	return &OrderController{
		Repo:     repo,
		Midtrans: midtrans,
		Valid:    valid,
	}
}

func (oc *OrderController) CreateOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpOrder vo.OrderRequest
		if err := c.Bind(&tmpOrder); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, view.BadRequest())
		}
		if err := oc.Valid.Struct(&tmpOrder); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, view.BadRequest())
		}
		user_id := middlewares.ExtractTokenUserId(c)

		order := entity.Order{
			Address:    tmpOrder.Address,
			ShipmentID: tmpOrder.ShipmentID,
			Status:     "wating",
			UserID:     uint(user_id),
		}

		res, err := oc.Repo.CreateOrder(order)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		// NumOrder := fmt.Sprintf("AE - %d", res.ID)
		// SnapMidtrans := oc.Midtrans.CreatePayout(NumOrder, res.TotalPay)
		// if SnapMidtrans == nil {
		// 	log.Warn("Error Snap")
		// 	return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		// }

		// return c.JSON(http.StatusCreated, vo.StatusCreated(NumOrder, SnapMidtrans))
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code": http.StatusOK,
			"data": res.Status,
		})
	}
}

func (oc *OrderController) CancelOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		err := oc.Repo.CancelOrder(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, vo.StatusCancelOrder())
	}
}

func (oc *OrderController) Payment() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderID := c.Param("order_id")
		UserID := middlewares.ExtractTokenUserId(c)
		result, err := oc.Repo.PayOrder(uint(UserID), OrderID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		response := vo.ResponseOrder{
			ID:       result.ID,
			Address:  result.Address,
			TotalPay: result.TotalPay,
			Status:   result.Status,
		}
		return c.JSON(http.StatusOK, vo.StatusPayOrder(response))
	}
}
