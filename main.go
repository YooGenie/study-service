package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"menu-service/config"
	"menu-service/config/handler"
	"menu-service/controller"
	"net/http"
	"os"
)

func main() {
	config.ConfigureEnvironment("./",  "STUDY_GENIE_DB_PASSWORD")
	xormDb := config.ConfigureDatabase()
	config.ConfigureLogger()
	// 이부분이 없으니까 500번 에러가 뜬다. 이유 찾아보기!
	defer func() {
		if r := recover(); r != nil {
			log.Errorln("Panic: %v", r)
			os.Exit(1)
		}
	}()
	e := config.ConfigureEcho()

	e.GET("/", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	e.Use(handler.CreateDatabaseContext(xormDb))
	//e.Use(handler.InitUserClaimMiddleware())
	//e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	controller.MenuController{}.Init(e.Group("/api/menu"))
	controller.StoreController{}.Init(e.Group("/api/store"))

	log.Info("study Service Server Started: Port=" + config.Config.HttpPort)
	e.Start(":" + config.Config.HttpPort)
}