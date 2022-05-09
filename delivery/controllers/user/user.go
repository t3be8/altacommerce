package user

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	view "github.com/t3be8/altacommerce/delivery/views"
	"github.com/t3be8/altacommerce/delivery/views/user"
	"github.com/t3be8/altacommerce/entity"
	userRepo "github.com/t3be8/altacommerce/repository/user"
	"github.com/t3be8/altacommerce/utils"
)

type UserController struct {
	Repo  userRepo.IUser
	Valid *validator.Validate
}

func New(repo userRepo.IUser, valid *validator.Validate) *UserController {
	return &UserController{
		Repo:  repo,
		Valid: valid,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpUser user.RegisterRequest

		if err := c.Bind(&tmpUser); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusBadRequest, user.BadRequest())
		}

		if err := uc.Valid.Struct(tmpUser); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusBadRequest, user.BadRequest())
		}

		pwd := tmpUser.Password
		hash, _ := utils.HashPassword(pwd)

		newUser := entity.User{
			Name:     tmpUser.Name,
			Email:    tmpUser.Email,
			Phone:    &tmpUser.Phone,
			Password: hash,
		}
		res, err := uc.Repo.Register(newUser)
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("berhasil register")
		return c.JSON(http.StatusCreated, user.SuccessInsert(res))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := user.LoginRequest{}

		if err := c.Bind(&param); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusBadRequest, user.BadRequest())
		}

		if err := uc.Valid.Struct(param); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusBadRequest, user.BadRequest())
		}

		data, match, err := uc.Repo.IsLogin(param.Email, param.Password)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"messages": err.Error(),
			})
		}

		if !match {
			return echo.ErrUnauthorized
		}

		res := user.LoginResponse{Data: data}

		if res.Token == "" {
			token, _ := CreateToken(int(data.ID))
			res.Token = token
			return c.JSON(http.StatusOK, view.OK(res, "Berhasil login"))
		}

		return c.JSON(http.StatusOK, view.OK(res, "Berhasil login"))
	}
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["expired"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("RU$SI4"))
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
