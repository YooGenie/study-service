package handler

import (
	"fmt"
	errors2 "menu-service/common/errors"
	"menu-service/config"
	"menu-service/dto/request"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

const (
	HeaderXRequestID    = "X-Request-ID"
	HeaderXActionID     = "X-Action-ID"
	HeaderXForwardedFor = "X-Forwarded-For"
	HeaderXRealIP       = "X-Real-IP"
	HeaderContentLength = "Content-Length"
)

type EchoRouter struct {
	once   sync.Once
	routes map[string]string
}

var echoRouter EchoRouter

func HandleErrorResponse(err error) (code int, message interface{}) {
	if ve, ok := err.(validator.ValidationErrors); ok {
		errorMessage := make([]string, 0)
		for _, e := range ve {
			errorType := e.Tag()
			if e.Param() != "" {
				errorType += "/"
				errorType += e.Param()
			}
			errorMessage = append(errorMessage, fmt.Sprintf("%s:%s", e.StructNamespace(), errorType))
		}
		code = http.StatusBadRequest
		message = errors2.ErrorResponseWrapper{Code: 10007, Message: strings.Join(errorMessage, "\n")}
	} else if ae, ok := err.(*errors2.ApiError); ok {
		code = ae.Code
		message = errors2.ErrorResponseWrapper{Code: ae.ErrorCode, Message: ae.Error()}
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else {
		code = http.StatusInternalServerError
	}
	if _, ok := message.(error); ok {
		var wrapped = make(map[string]interface{})
		wrapped["message"] = message.(error).Error()
		message = wrapped
	} else if _, ok := message.(string); ok {
		var wrapped = make(map[string]interface{})
		wrapped["message"] = message
		message = wrapped
	}

	return
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code, message := HandleErrorResponse(err)

	stack := make([]byte, 2048)
	length := runtime.Stack(stack, true)
	stackTrace := fmt.Sprintf("[ERROR] %v %s\n", err, stack[:length])

	req := c.Request()
	newErrorLog := NewErrorLog(req)
	for _, name := range c.ParamNames() {
		newErrorLog.Params[name] = c.Param(name)
	}
	res := c.Response()
	newErrorLog.Status = code
	newErrorLog.BytesSent = res.Size
	newErrorLog.Controller, newErrorLog.Action = echoRouter.getControllerAndAction(c)

	newErrorLog.Code = strconv.Itoa(code)
	newErrorLog.Error = err.Error()
	newErrorLog.StackTrace = stackTrace

	hostname, err := os.Hostname()
	// logrus.WithError(err).Error("Fail to get hostname")
	logrus.Infoln("hostname:", hostname)
	newErrorLog.Hostname = hostname

	//service.ErrorLogService().Create(c.Request().Context(), newErrorLog)

	c.JSON(code, message)
}

func NewErrorLog(req *http.Request) dto.ErrorLog {
	serviceName := config.Config.Service.Name
	realIP := req.RemoteAddr
	if ip := req.Header.Get(HeaderXForwardedFor); ip != "" {
		realIP = strings.Split(ip, ", ")[0]
	} else if ip := req.Header.Get(HeaderXRealIP); ip != "" {
		realIP = ip
	} else {
		realIP, _, _ = net.SplitHostPort(realIP)
	}

	path := req.URL.Path
	if path == "" {
		path = "/"
	}

	requestLength, _ := strconv.ParseInt(req.Header.Get(HeaderContentLength), 10, 64)

	params := map[string]interface{}{}
	for k, v := range req.URL.Query() {
		params[k] = v[0]
	}

	c := dto.ErrorLog{
		Service:       serviceName,
		Timestamp:     time.Now(),
		RemoteIP:      realIP,
		Host:          req.Host,
		Uri:           req.RequestURI,
		Method:        req.Method,
		Path:          path,
		Params:        params,
		Referer:       req.Referer(),
		UserAgent:     req.UserAgent(),
		RequestLength: requestLength,
		AuthToken:     req.Header.Get(echo.HeaderAuthorization),
	}

	return c
}

func (er *EchoRouter) getControllerAndAction(c echo.Context) (controller, action string) {
	er.once.Do(func() { er.initialize(c) })

	if v := c.Get("controller"); v != nil {
		if controllerName, ok := v.(string); ok {
			controller = controllerName
		}
	}
	if v := c.Get("action"); v != nil {
		if actionName, ok := v.(string); ok {
			action = actionName
		}
	}

	if controller == "" || action == "" {
		handlerName := er.routes[fmt.Sprintf("%s+%s", c.Path(), c.Request().Method)]
		controller, action = er.convertHandlerNameToControllerAndAction(handlerName)
	}
	return
}

func (er *EchoRouter) initialize(c echo.Context) {
	er.routes = make(map[string]string)
	for _, r := range c.Echo().Routes() {
		path := r.Path
		if len(path) == 0 || path[0] != '/' {
			path = "/" + path
		}
		er.routes[fmt.Sprintf("%s+%s", path, r.Method)] = r.Name
	}
}

func (er EchoRouter) convertHandlerNameToControllerAndAction(handlerName string) (controller, action string) {
	handlerSplitIndex := strings.LastIndex(handlerName, ".")
	if handlerSplitIndex == -1 || handlerSplitIndex >= len(handlerName) {
		controller, action = "", handlerName
	} else {
		controller, action = handlerName[:handlerSplitIndex], handlerName[handlerSplitIndex+1:]
	}

	// 1. find this pattern: "(controller)"
	controller = controller[strings.Index(controller, "(")+1:]
	if index := strings.Index(controller, ")"); index > 0 {
		controller = controller[:index]
	}
	// 2. remove pointer symbol
	controller = strings.TrimPrefix(controller, "*")
	// 3. split by "/"
	if index := strings.LastIndex(controller, "/"); index > 0 {
		controller = controller[index+1:]
	}

	// remove function symbol
	action = strings.TrimRight(action, ")-fm")
	return
}
