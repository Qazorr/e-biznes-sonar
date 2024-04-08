package controllers

import (
	"exercise_4/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var categoryErrors = map[string]string{
	"categoryNotFound": "Category not found",
}

type CategoryController struct {
	DB *gorm.DB
}

func (cc *CategoryController) GetAllCategories(c echo.Context) error {
	var categories []models.Category
	cc.DB.Find(&categories)
	return c.JSONPretty(http.StatusOK, categories, "  ")
}

func (cc *CategoryController) GetCategoryByID(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if err := cc.DB.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": categoryErrors["categoryNotFound"]})
	}
	return c.JSONPretty(http.StatusOK, category, "  ")
}

func (cc *CategoryController) CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return err
	}
	cc.DB.Create(category)
	return c.JSONPretty(http.StatusOK, category, "  ")
}

func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if err := cc.DB.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": categoryErrors["categoryNotFound"]})
	}
	if err := c.Bind(&category); err != nil {
		return err
	}
	cc.DB.Save(&category)
	return c.JSONPretty(http.StatusOK, category, "  ")
}

func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if err := cc.DB.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": categoryErrors["categoryNotFound"]})
	}
	cc.DB.Delete(&category)
	return c.NoContent(http.StatusNoContent)
}
