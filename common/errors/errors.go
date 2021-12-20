package errors

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

var (
	ErrAuthentication       = new(http.StatusUnauthorized, 10001, "로그인이 필요합니다.")
	ErrAuthorization        = new(http.StatusForbidden, 10002, "잘못된 접근입니다.")
	ErrNoResult             = new(http.StatusNotFound, 10003, "결과를 찾을 수 없습니다.")
	ErrDuplicationRegister  = new(http.StatusBadRequest, 10004, "이미 등록되었습니다.")
	ErrNoDeletable          = new(http.StatusGone, 10005, "삭제할 수 없는 데이터입니다.")
	ErrProcessingImpossible = new(http.StatusGone, 10006, "처리할 수 없는 상태입니다.")
	// 10007 : go-validator Validation 오류
	ErrNotValid       = new(http.StatusBadRequest, 10008, "잘못된 요청입니다.")
	ErrIdInconsistent = new(http.StatusInternalServerError, 90000, "ID가 일치하지 않습니다.")
)

type ApiError struct {
	*echo.HTTPError
	ErrorCode int
}

type ErrorResponseWrapper struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

func new(statusCode int, code int, message interface{}) error {
	return &ApiError{ErrorCode: code, HTTPError: echo.NewHTTPError(statusCode, message)}
}

func ValidationError(message string) error {
	return new(http.StatusBadRequest, 10008, message)
}

func Throw(cause error) error {
	return &ApiError{HTTPError: echo.NewHTTPError(http.StatusInternalServerError, cause)}
}

func ApiParamValidError(err error) error {
	return new(http.StatusBadRequest, 90001, err.Error())
}

