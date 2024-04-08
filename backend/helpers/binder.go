package helpers

import (
	"exercise_4/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BindProductController(e *echo.Echo, db *gorm.DB) {
	productController := &controllers.ProductController{DB: db}

	e.GET("/products", productController.GetAllProducts)
	e.GET("/products/:id", productController.GetProductByID)
	e.POST("/products", productController.CreateProduct)
	e.PUT("/products/:id", productController.UpdateProduct)
	e.DELETE("/products/:id", productController.DeleteProduct)
}

func BindCartItemController(e *echo.Echo, db *gorm.DB) {
	cartItemController := &controllers.CartItemController{DB: db}

	e.GET("/cart", cartItemController.GetAllCartItems)
	e.GET("/cart/:id", cartItemController.GetCartItemByID)
	e.POST("/cart", cartItemController.CreateCartItem)
	e.PUT("/cart/:id", cartItemController.UpdateCartItem)
	e.DELETE("/cart/:id", cartItemController.DeleteCartItem)
}

func BindCategoryController(e *echo.Echo, db *gorm.DB) {
	categoryController := &controllers.CategoryController{DB: db}

	e.GET("/categories", categoryController.GetAllCategories)
	e.GET("/categories/:id", categoryController.GetCategoryByID)
	e.POST("/categories", categoryController.CreateCategory)
	e.PUT("/categories/:id", categoryController.UpdateCategory)
	e.DELETE("/categories/:id", categoryController.DeleteCategory)
}

func BindPaymentController(e *echo.Echo, db *gorm.DB) {
	paymentController := &controllers.PaymentController{DB: db}

	e.POST("/payments", paymentController.SendPayment)
}
