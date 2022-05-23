package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)
// API 콜할 때마다 공통으로 콜하는 부분이 있다. 이걸 매번 API 콜할 때마다 접근해야하는 불편함을 없애기 위해서 미들웨어를 사용한다.
// 미들웨어에는 공통으로 작동하는 것들을 넣는다.

// 에코설정하기
func ConfigureEcho() *echo.Echo {
	e := echo.New()
	e.Validator = RegisterValidator()
	e.HideBanner = true


	//e.Pre()는 라우팅 하기전 특정 API 불러서 컨트롤러 하기전에 Pre 먼저 자동하는 것
	e.Pre(middleware.RemoveTrailingSlash()) // URL 끝에 / 넣어서 오면 슬래시를 자동으로 빼주는 것이다. 이건 함수 타기전에 해야하는 일이다. use로 시작하는 것은 컨트롤러(헨들러) 시작했을 때 해주는것이다
	e.Use(middleware.Recover()) // 죽었을 때 다시 살리는 것
	e.Use(middleware.CORS()) //  CORS 셋팅해주는 것이다. 이때 메소드를 정할 수 있다. 뭐뭐 메소드로 바뀐다. 예를 들으면 우리는 조회만 하니까 GET밖에 사용안한다.  post만 받는다. 이런 식으로 메소드를 정할 수 있다.

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://", "https://", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.RequestID()) //요청 ID 미들웨어는 요청에 대한 고유 ID를 생성합니다.
	//e.Use(middleware.JWTWithConfig()) //JWT(JSON Web Token) 인증 미들웨어를 제공
	
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