package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ConfigureEcho() *echo.Echo {
	e := echo.New()
	e.Validator = RegisterValidator()
	e.HideBanner = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	return e
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) (err error) {
	return cv.validator.Struct(i)
}

func RegisterValidator() *CustomValidator {
	customValidator := validator.New()

	return &CustomValidator{validator: customValidator}
}