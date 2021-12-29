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
}

