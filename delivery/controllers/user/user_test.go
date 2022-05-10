package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"github.com/t3be8/altacommerce/entity"
	"gorm.io/gorm"
)

var token string

func TestLogin(t *testing.T) {
	e := echo.New()

	requestBody, _ := json.Marshal(map[string]interface{}{
		"email":    "jdoe@test.com",
		"password": "admin234",
	})

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	context := e.NewContext(req, res)
	context.SetPath("/login")

	controller := New(&mockRepoUser{}, validator.New())
	controller.Login()(context)

	var response ResponseStructure

	json.Unmarshal([]byte(res.Body.Bytes()), &response)
	assert.Equal(t, 200, response.Code)
	assert.True(t, response.Status)
	assert.NotNil(t, response.Data)
	data := response.Data.(map[string]interface{})
	token = data["Token"].(string)
}

func TestRegister(t *testing.T) {
	t.Run("Success Register", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"name":     "John Doe",
			"email":    "jdoe@test.com",
			"phone":    "08967544321",
			"password": "admin234",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")

		registerController := New(&mockRepoUser{}, validator.New())
		registerController.Register()(context)

		var response ResponseStructure

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "John Doe", response.Data.(map[string]interface{})["Name"])
		assert.True(t, response.Status)
		assert.Equal(t, 201, response.Code)
	})

	t.Run("Error Validasi", func(t *testing.T) {
		e := echo.New()
		rb, _ := json.Marshal(map[string]interface{}{
			"email": "asdsefs",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(rb)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")

		registerController := New(&erorrMockUserRepository{}, validator.New())
		registerController.Register()(context)

		var response ResponseStructure

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Warn(response)
		assert.False(t, response.Status)
		assert.Nil(t, response.Data)
		assert.Equal(t, 400, response.Code)
	})

	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"email": "jdoe@test.com",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")

		registerController := New(&erorrMockUserRepository{}, validator.New())
		registerController.Register()(context)

		var response ResponseStructure

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Warn(response)
		assert.False(t, response.Status)
		assert.Nil(t, response.Data)
		assert.Equal(t, 400, response.Code)
	})
	t.Run("Error Insert DB", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":  "Galang",
			"email": "gadipuran@test.col",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/register")

		registerController := New(&erorrMockUserRepository{}, validator.New())
		registerController.Register()(context)

		var response ResponseStructure

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.False(t, response.Status)
		assert.Nil(t, response.Data)
		assert.Equal(t, 500, response.Code)
	})

}

type ResponseStructure struct {
	Code    int
	Message string
	Status  bool
	Data    interface{}
}

type mockRepoUser struct{}

func (mru *mockRepoUser) Register(newUser entity.User) (entity.User, error) {
	return newUser, nil
}

func (mru *mockRepoUser) IsLogin(email, password string) (entity.User, bool, error) {
	return entity.User{Model: gorm.Model{ID: uint(1)}, Name: "John Doe", Email: "jdoe@test.com"}, true, nil
}

type erorrMockUserRepository struct{}

func (emur *erorrMockUserRepository) Register(newPegawai entity.User) (entity.User, error) {
	return entity.User{}, errors.New("tidak bisa insert data")
}
func (emur *erorrMockUserRepository) IsLogin(email, password string) (entity.User, bool, error) {
	return entity.User{}, false, errors.New("tidak bisa select data")
}
