package controller

import (
	"net/http"
	"study-service/common/errors"
	requestDto "study-service/dto/request"
	service2 "study-service/service"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type ClickController struct {
}

func (controller ClickController) Init(g *echo.Group) {
	g.POST("", controller.Create)
}

func (ClickController) Create(ctx echo.Context) error {
	var clickCreate = requestDto.ClickCreate{}

	if err := ctx.Bind(&clickCreate); err != nil {
		return errors.ApiParamValidError(err)
	}

	if err := clickCreate.Validate(ctx); err != nil {
		log.Errorf("Create Error:  %s", err.Error())
		return err
	}

	err := service2.ClickService().Create(ctx.Request().Context(), clickCreate)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
