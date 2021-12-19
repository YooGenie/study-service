package controller

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"menu-service/common/errors"
	requestDto "menu-service/dto/request"
	"menu-service/store/service"
	"net/http"
)

type StoreController struct {
}

func (controller StoreController) Init(g *echo.Group) {
	g.POST("", controller.Create)
}

func (StoreController) Create(ctx echo.Context) error {
	var storeCreate = requestDto.StoreCreate{}

	if err := ctx.Bind(&storeCreate); err != nil {
		return errors.ApiParamValidError(err)
	}

	if err := storeCreate.Validate(ctx); err != nil {
		log.Errorf("Create Error:  %s", err.Error())
		return err
	}

	//서비스부분 연결
	err := service.StoreService().Create(ctx.Request().Context(), storeCreate)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}


