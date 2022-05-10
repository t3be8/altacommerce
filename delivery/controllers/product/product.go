package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	view "github.com/t3be8/altacommerce/delivery/views"

	productView "github.com/t3be8/altacommerce/delivery/views/product"
	"github.com/t3be8/altacommerce/entity"
	productRepo "github.com/t3be8/altacommerce/repository/product"
)

type ProductController struct {
	Repo  productRepo.IProduct
	Valid *validator.Validate
}

func New(repo productRepo.IProduct, valid *validator.Validate) *ProductController {
	return &ProductController{
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

func (pc *ProductController) InsertProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpProduct productView.InsertProductRequest

		if err := c.Bind(&tmpProduct); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusBadRequest, productView.BadRequest())
		}

		if err := pc.Valid.Struct(tmpProduct); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusBadRequest, productView.BadRequest())
		}

		res, err := pc.Repo.InsertProduct(entity.Product{Name: tmpProduct.Name, Description: tmpProduct.Description, Price: tmpProduct.Price,
			Stok: tmpProduct.Stok, Image: tmpProduct.Image})

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("berhasil insert product")
		return c.JSON(http.StatusCreated, productView.SuccessInsert(res))
	}
}

func (pc *ProductController) SelectProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := ExtractTokenUserId(c)

		fmt.Println(id)
		res, err := pc.Repo.SelectProduct()

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("berhasil select product")
		return c.JSON(http.StatusOK, productView.SuccessSelect(res))
	}
}

func (pc *ProductController) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpProduct productView.UpdateProductRequest
		id := ExtractTokenUserId(c)

		fmt.Println(id)

		res, err := pc.Repo.UpdateProduct(tmpProduct.ID, tmpProduct.Stok)
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		log.Info("berhasil update stock pada product")
		return c.JSON(http.StatusCreated, productView.SuccessUpdate(res))
	}
}

func (pc *ProductController) DeletedProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpProduct productView.DeleteProductRequest
		id := ExtractTokenUserId(c)

		fmt.Println(id)
		res, err := pc.Repo.DeletedProduct(tmpProduct.ID)
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		log.Info("berhasi delete product")
		return c.JSON(http.StatusOK, productView.SuccessDelete(res))
	}
}

func (pc *ProductController) GetAllProductById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := ExtractTokenUserId(c)

		IdProduct := c.Param("id")
		conID, _ := strconv.Atoi(IdProduct)

		fmt.Println(id)
		res, err := pc.Repo.GetAllById(conID)

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("berhasil select product by ID")
		return c.JSON(http.StatusOK, productView.SuccessSelect(res))
	}
}

func (pc *ProductController) GetAllProductByCategory() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := ExtractTokenUserId(c)

		idCategory := c.Param("produccategoryid")
		conID, _ := strconv.Atoi(idCategory)

		fmt.Println(id)
		res, err := pc.Repo.GetAllByCategory(conID)

		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		log.Info("berhasil select product by Category")
		return c.JSON(http.StatusOK, productView.SuccessSelect(res))
	}
}
