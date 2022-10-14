package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"study-service/config"
	"study-service/config/handler"
	"study-service/controller"
)

func main() {
	config.ConfigureEnvironment("./", "STUDY_GENIE_DB_PASSWORD", "STUDY_GENIE_ENCRYPT_KEY", "KAKAO_API_KEY") //환경변수 설정

	//DB 트랜잭션을 처리하기 위해서 하나의 트랙잭션을 관리한다. API를 콜할 때마다 DB를 연결하면 트랙잭션을 처리할 수 없다.
	//처음 한번만 사용한다 그래서 미들웨어에서 사용할 수 있도록 구현해야한다.
	xormDb := config.ConfigureDatabase() //DB 설정

	config.ConfigureLogger() // log 설정
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
	controller.AuthController{}.Init(e.Group("/api/auth"))
	controller.MemberController{}.Init(e.Group("/api/member"))
	controller.ClickController{}.Init(e.Group("/api/click"))
	controller.PdfController{}.Init(e.Group("/api/pdf"))

	log.Info("study Service Server Started: Port=" + config.Config.HttpPort)
	e.Start(":" + config.Config.HttpPort)
}
