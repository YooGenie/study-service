package main

import (
	"net/http"
	"os"
	"study-service/config"
	"study-service/config/handler"
	"study-service/controller"

	_ "study-service/docs"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	config.ConfigureEnvironment("./", "STUDY_GENIE_DB_PASSWORD", "STUDY_GENIE_ENCRYPT_KEY", "KAKAO_API_KEY") //환경변수 설정

	//DB 트랜잭션을 처리하기 위해서 하나의 트랙잭션을 관리한다. API를 콜할 때마다 DB를 연결하면 트랙잭션을 처리할 수 없다.
	//처음 한번만 사용한다 그래서 미들웨어에서 사용할 수 있도록 구현해야한다.
	xormDb := config.ConfigureDatabase() //DB 설정

	//config.ConfigureLogger() // log 설정
	// 이부분이 없으니까 500번 에러가 뜬다. 이유 찾아보기!
	defer func() {
		if r := recover(); r != nil {
			log.Errorln("Panic: %v", r)
			os.Exit(1)
		}
	}()

	//에코함수 불러오기
	e := config.ConfigureEcho()

	e.GET("/", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(handler.CreateDatabaseContext(xormDb)) // 뭐를 부르던지 무조건 와서 실행해서 데이터 베이스 컨텍스트를 만들어 주는 것이다. DB를 연결 해주는 거니까 모두 필요하다.
	//e.Use(handler.InitUserClaimMiddleware()) // 사용자 정보 토큰 만들어 주는 것 이것도 모든 API에 다 사용하는데 각각에 적어주는 것보다 미들웨어안에 적어 넣어으면 자동으로 실행 된다.
	//e.HTTPErrorHandler = handler.CustomHTTPErrorHandler // 메시지와 관리를 하기 위해서 에러핸드링 해주는 것이다. 에러가 떨어질 때 자동으로 에러 핸들러를 불러서 에러를 컨트롤 해준다.

	controller.MenuController{}.Init(e.Group("/api/menu"))
	controller.StoreController{}.Init(e.Group("/api/store"))
	controller.AuthController{}.Init(e.Group("/api/auth"))
	controller.MemberController{}.Init(e.Group("/api/member"))
	controller.ClickController{}.Init(e.Group("/api/click"))
	controller.PdfController{}.Init(e.Group("/api/pdf"))
	controller.EmailController{}.Init(e.Group("/api/email"))

	log.Info("study Service Server Started: Port=" + config.Config.HttpPort)
	e.Start(":" + config.Config.HttpPort)

	// 핫픽스 테스트

	// 다른줄에 데브 브랜치

}
