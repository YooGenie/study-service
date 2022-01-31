package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"study-service/config"
	"study-service/config/handler"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

var (
	echoApp *echo.Echo
	xormDb  config.DatabaseWrapper
)

func init() {
	runtime.GOMAXPROCS(1)
	os.Setenv("CONFIGOR_ENV", "test")
	config.ConfigureEnvironment("../", "STUDY_GENIE_ENCRYPT_KEY")
	xormDb = config.ConfigureDatabase()
	config.ConfigureLogger()

	e := config.ConfigureEcho()

	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	echoApp = e
}

type EchoContextBuilder struct {
	request     *http.Request
	paramKeys   []string
	paramValues []string
	context     map[string]interface{}
}

func NewRequest(req *http.Request) EchoContextBuilder {
	return EchoContextBuilder{request: req, paramKeys: make([]string, 0), paramValues: make([]string, 0), context: make(map[string]interface{})}
}

func (cb EchoContextBuilder) WithContext(key string, value interface{}) EchoContextBuilder {
	cb.context[key] = value
	return cb
}



func (cb EchoContextBuilder) WithParam(key string, value string) EchoContextBuilder {
	cb.paramKeys = append(cb.paramKeys, key)
	cb.paramValues = append(cb.paramValues, value)
	return cb
}

func (cb EchoContextBuilder) Handle(handlerFunc echo.HandlerFunc) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	echo := echoApp.NewContext(cb.request, rec)

	echo.SetParamNames(cb.paramKeys...)
	echo.SetParamValues(cb.paramValues...)

	// DB Context는 기본 설정으로 간주.
	ctx := context.WithValue(echo.Request().Context(), config.ContextDBKey, xormDb)

	for k, v := range cb.context {
		ctx = context.WithValue(echo.Request().Context(), k, v)
	}

	echo.SetRequest(echo.Request().WithContext(ctx))

	session := handler.CreateDatabaseContext(xormDb)
	err := session(handlerFunc)(echo)
	if err != nil {
		echo.Error(err)
	}
	return rec
}

func jsonResponse(rec *httptest.ResponseRecorder) (result map[string]interface{}, err error) {
	str := rec.Body.String()
	result = make(map[string]interface{})
	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		return
	}
	return
}
