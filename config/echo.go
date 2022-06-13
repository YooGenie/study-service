package config

import (
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
