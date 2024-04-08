package controllers

import (
	"exercise_4/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var productErrors = map[string]string{
	"productNotFound": "Product not found",
}

type ProductController struct {
	DB *gorm.DB
}

func (pc *ProductController) GetAllProducts(c echo.Context) error {
	var products []models.Product
	pc.DB.Find(&products)
	return c.JSONPretty(http.StatusOK, products, "  ")
}

func (pc *ProductController) GetProductByID(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": productErrors["productNotFound"]})
	}
	return c.JSONPretty(http.StatusOK, product, "  ")
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	pc.DB.Create(product)
	return c.JSONPretty(http.StatusOK, product, "  ")
}

func (pc *ProductController) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": productErrors["productNotFound"]})
	}
	if err := c.Bind(&product); err != nil {
		return err
	}
	pc.DB.Save(&product)
	return c.JSONPretty(http.StatusOK, product, "  ")
}

func (pc *ProductController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": productErrors["productNotFound"]})
	}
	pc.DB.Delete(&product)
	return c.NoContent(http.StatusNoContent)
}
