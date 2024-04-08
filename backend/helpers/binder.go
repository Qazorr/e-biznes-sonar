package helpers

import (
	"exercise_4/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BindProductController(e *echo.Echo, db *gorm.DB) {
	productController := &controllers.ProductController{DB: db}
	const productRoute = "/products"
	const productIDRoute = productRoute + "/:id"

	e.GET(productRoute, productController.GetAllProducts)
	e.GET(productIDRoute, productController.GetProductByID)
	e.POST(productRoute, productController.CreateProduct)
	e.PUT(productIDRoute, productController.UpdateProduct)
	e.DELETE(productIDRoute, productController.DeleteProduct)
}

func BindCartItemController(e *echo.Echo, db *gorm.DB) {
	cartItemController := &controllers.CartItemController{DB: db}
	const cartItemRoute = "/cart"
	const cartItemIDRoute = cartItemRoute + "/:id"

	e.GET(cartItemRoute, cartItemController.GetAllCartItems)
	e.GET(cartItemIDRoute, cartItemController.GetCartItemByID)
	e.POST(cartItemRoute, cartItemController.CreateCartItem)
	e.PUT(cartItemIDRoute, cartItemController.UpdateCartItem)
	e.DELETE(cartItemIDRoute, cartItemController.DeleteCartItem)
}

func BindCategoryController(e *echo.Echo, db *gorm.DB) {
	categoryController := &controllers.CategoryController{DB: db}
	const categoryRoute = "/categories"
	const categoryIDRoute = categoryRoute + "/:id"

	e.GET(categoryRoute, categoryController.GetAllCategories)
	e.GET(categoryIDRoute, categoryController.GetCategoryByID)
	e.POST(categoryRoute, categoryController.CreateCategory)
	e.PUT(categoryIDRoute, categoryController.UpdateCategory)
	e.DELETE(categoryIDRoute, categoryController.DeleteCategory)
}

func BindPaymentController(e *echo.Echo, db *gorm.DB) {
	paymentController := &controllers.PaymentController{DB: db}

	e.POST("/payments", paymentController.SendPayment)
}
