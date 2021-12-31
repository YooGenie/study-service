package controller

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	responseDto "menu-service/dto/response"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMenuController_Create(t *testing.T) {
	DatabaseFixture{}.setUpDefault(xormDb.Engine)

	t.Run("Create Menu", func(t *testing.T) {
		// given
		requestBody := `{
          "name": "육개장",
            "price": 9000,
            "description": "국물이 진국이에요"
    }`
		req := httptest.NewRequest(echo.POST, fmt.Sprintf("/api/menu"), strings.NewReader(requestBody))
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

		// when
		rec := NewRequest(req).
			Handle(MenuController{}.Create)
		result := responseDto.ApiError{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		// then
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("Create Menu no price", func(t *testing.T) {
		// given
		requestBody := `{
          "name": "떡볶이",
            "price": 0,
            "description": "국물이 진국이에요"
    }`
		req := httptest.NewRequest(echo.POST, fmt.Sprintf("/api/menu"), strings.NewReader(requestBody))
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

		// when
		rec := NewRequest(req).
			Handle(MenuController{}.Create)
		result := responseDto.ApiError{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		// then
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestMenuController_Update(t *testing.T) {
	DatabaseFixture{}.setUpDefault(xormDb.Engine)

	t.Run("Update Menu", func(t *testing.T) {
		// given
		requestBody := `{
			"id" : 3,
            "name": "신라면",
            "price": 3000,
            "description": "국물이 진국이에요"
    }`
		req := httptest.NewRequest(echo.PUT, fmt.Sprintf("/api/menu"), strings.NewReader(requestBody))
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

		// when
		rec := NewRequest(req).
			Handle(MenuController{}.Update)
		result := responseDto.ApiError{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		// then
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("Update Menu no id", func(t *testing.T) {
		// given
		requestBody := `{
			"id"=5,
            "name": "순대",
            "price": 3500,
            "description": "순대 맛있다"
    }`
		req := httptest.NewRequest(echo.PUT, fmt.Sprintf("/api/menu"), strings.NewReader(requestBody))
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

		// when
		rec := NewRequest(req).
			Handle(MenuController{}.Update)
		result := responseDto.ApiError{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		// then
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
