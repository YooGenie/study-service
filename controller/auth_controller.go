package controller

import (
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

	return ctx.JSON(http.StatusOK, "")
}
