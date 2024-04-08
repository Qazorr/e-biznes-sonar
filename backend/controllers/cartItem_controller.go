package controllers

import (
	"exercise_4/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CartItemController struct {
	DB *gorm.DB
}

func (cic *CartItemController) GetAllCartItems(c echo.Context) error {
	var cartItems []models.CartItem
	cic.DB.Find(&cartItems)
	return c.JSONPretty(http.StatusOK, cartItems, "  ")
}

func (cic *CartItemController) GetCartItemByID(c echo.Context) error {
	id := c.Param("id")
	var cartItem models.CartItem
	if err := cic.DB.First(&cartItem, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "CartItem not found"})
	}
	return c.JSONPretty(http.StatusOK, cartItem, "  ")
}

func (cic *CartItemController) CreateCartItem(c echo.Context) error {
	cartItem := new(models.CartItem)
	if err := c.Bind(cartItem); err != nil {
		return err
	}

	// if is present, increase quantity otherwise just create it
	existingCartItem := new(models.CartItem)
	if err := cic.DB.Where("product_id = ?", cartItem.ProductID).First(existingCartItem).Error; err != nil {
		cic.DB.Create(cartItem)
	} else {
		existingCartItem.Quantity++
		cic.DB.Save(existingCartItem)
	}

	return c.JSONPretty(http.StatusOK, cartItem, "  ")
}

func (cic *CartItemController) UpdateCartItem(c echo.Context) error {
	id := c.Param("id")
	var cartItem models.CartItem
	if err := cic.DB.First(&cartItem, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "CartItem not found"})
	}
	if err := c.Bind(&cartItem); err != nil {
		return err
	}
	cic.DB.Save(&cartItem)
	return c.JSONPretty(http.StatusOK, cartItem, "  ")
}

func (cic *CartItemController) DeleteCartItem(c echo.Context) error {
	id := c.Param("id")
	var cartItem models.CartItem
	if err := cic.DB.First(&cartItem, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "CartItem not found"})
	}

	// reversed to adding - if quantity is 0, delete it else decrement
	if cartItem.Quantity > 0 {
		cartItem.Quantity--
		if cartItem.Quantity == 0 {
			cic.DB.Delete(&cartItem)
		} else {
			cic.DB.Save(&cartItem)
		}
	}

	return c.NoContent(http.StatusNoContent)
}
