package controller

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"menu-service/common/errors"
	requestDto "menu-service/dto/request"
	"menu-service/store/service"
	"net/http"
	"strconv"
)

type StoreController struct {
}

func (controller StoreController) Init(g *echo.Group) {
	g.POST("", controller.Create)
	g.GET("/:no", controller.GetStoreById)
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

	err := service.StoreService().Create(ctx.Request().Context(), storeCreate)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (StoreController) GetStoreById(ctx echo.Context) error {
	storeNo, err := strconv.ParseInt(ctx.Param("no"), 10, 64)
	if err != nil {
		return errors.ApiParamValidError(err)
	}

	menu, _ := service.StoreService().GetStoreById(ctx.Request().Context(), storeNo)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, menu)

}
