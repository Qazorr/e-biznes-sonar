package controllers

import (
	"exercise_4/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PaymentController struct {
	DB *gorm.DB
}

func (pc *PaymentController) SendPayment(c echo.Context) error {
	var cartItems []models.CartItem
	pc.DB.Where("deleted_at IS NULL").Delete(&cartItems)
	return c.JSONPretty(http.StatusOK, "Payment sent", "  ")
}
