package controller

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"menu-service/common/errors"
	requestDto "menu-service/dto/request"
	"menu-service/member/service"
	"net/http"
)

type MemberController struct {
}

func (controller MemberController) Init(g *echo.Group) {
	g.POST("", controller.Create)
}

func (MemberController) Create(ctx echo.Context) error {
	var memberCreate = requestDto.MemberCreate{}

	if err := ctx.Bind(&memberCreate); err != nil {
		return errors.ApiParamValidError(err)
	}

	if err := memberCreate.Validate(ctx); err != nil {
		log.Errorf("Create Error:  %s", err.Error())
		return err
	}

	err := service.MemberService().Create(ctx.Request().Context(), memberCreate)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}
