package controller

import (
	"menu-service/auth/service"
	"menu-service/common/errors"
	requestDto "menu-service/dto/request"
	"net/http"

	"github.com/labstack/echo"
)

type AuthController struct {
}

func (controller AuthController) Init(g *echo.Group) {
	g.POST("/login", controller.AuthAdminWithEmailAndPassword)
}

func (AuthController) AuthAdminWithEmailAndPassword(ctx echo.Context) (err error) {
	var adminSignIn requestDto.AdminSignIn
	if err := ctx.Bind(&adminSignIn); err != nil {
		return errors.ApiParamValidError(err)
	}

	if err := adminSignIn.Validate(ctx); err != nil {
		return err
	}

	jwtToken, err := service.AuthService().AuthWithSignIdPassword(ctx.Request().Context(), adminSignIn)
	if err != nil {
		if err == errors.ErrAuthentication {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil || len(refreshToken.Value) == 0 {
		cookie := new(http.Cookie)
		cookie.Name = "refreshToken"
		cookie.Value = jwtToken.RefreshToken
		cookie.HttpOnly = true
		cookie.Path = "/"
		ctx.SetCookie(cookie)
	} else {
		refreshToken.Value = jwtToken.RefreshToken
		refreshToken.HttpOnly = true
		refreshToken.Path = "/"
		ctx.SetCookie(refreshToken)
	}

	result := map[string]string{}
	result["accessToken"] = jwtToken.AccessToken
	return ctx.JSON(http.StatusOK, result)

	return ctx.JSON(http.StatusOK, "")
}
