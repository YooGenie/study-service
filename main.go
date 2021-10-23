package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"menu-service/common"
	"menu-service/config"
	"menu-service/controllers"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"xorm.io/core"
)

var (
	xormDb *xorm.Engine
)

func init() {
	config.InitConfig("config/config.json")

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	dbPassword := os.Getenv("STUDY_GENIE_DB_PASSWORD")
	if len(dbPassword) == 0 {
		panic(fmt.Errorf("No STUDY_GENIE_DB_PASSWORD system env variable\n"))
	}

	dbConnection := fmt.Sprintf("%s:%s%s", config.Config.Database.User, dbPassword, config.Config.Database.Connection)
	db, err := xorm.NewEngine(config.Config.Database.Driver, dbConnection)
	if err != nil {
		panic(fmt.Errorf("Database open error: connection url: %s, error: %s \n", config.Config.Database.Connection, err))
	} else {
		fmt.Println("DB connected: ", config.Config.Database.Connection)
	}

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(10 * time.Minute)

	db.ShowSQL(true)
	db.Logger().SetLevel(core.LOG_INFO)
	xormDb = db
}

func main() {
	defer xormDb.Close()

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(common.InitContextDB(xormDb))

	controllers.MenuController{}.Init(e.Group("/api/menu"))

	log.Info("Study Service Server Started: Port=" + config.Config.HttpPort)
	e.Start(":" + config.Config.HttpPort)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type CustomValidator struct {
	validator *validator.Validate
}
