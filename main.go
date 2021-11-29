package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"menu-service/config"
	"menu-service/config/handler"
	"menu-service/controllers"
	"net/http"
)

func main() {
	config.ConfigureEnvironment("./",  "STUDY_GENIE_DB_PASSWORD")
	xormDb := config.ConfigureDatabase()
	config.ConfigureLogger()

	e := config.ConfigureEcho()

	e.GET("/", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	e.Use(handler.CreateDatabaseContext(xormDb))
	//e.Use(handler.InitUserClaimMiddleware())
	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	controllers.MenuController{}.Init(e.Group("/api/menu"))

	log.Info("Study Service Server Started: Port=" + config.Config.HttpPort)
	e.Start(":" + config.Config.HttpPort)
}