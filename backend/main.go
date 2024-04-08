package main

import (
	"exercise_4/helpers"
	"exercise_4/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Product{}, &models.Category{}, &models.CartItem{})

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	helpers.BindCartItemController(e, db)
	helpers.BindCategoryController(e, db)
	helpers.BindProductController(e, db)
	helpers.BindPaymentController(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
