package controller

import (
	"net/http"
	"strconv"
	"study-service/common/errors"
	requestDto "study-service/dto/request"
	service2 "study-service/service"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type StoreController struct {
}

func (controller StoreController) Init(g *echo.Group) {
	g.POST("", controller.Create)
	g.GET("/:no", controller.GetStoreById)
	g.PUT("/:no", controller.Update)
	g.GET("", controller.GetStores)
	g.PUT("/delete/:no", controller.Delete)
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

	err := service2.StoreService().Create(ctx.Request().Context(), storeCreate)
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

	menu, _ := service2.StoreService().GetStoreByNo(ctx.Request().Context(), storeNo)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, menu)

}

func (StoreController) Update(ctx echo.Context) error {
	storeNo, err := strconv.ParseInt(ctx.Param("no"), 10, 64)

	if err != nil {
		return errors.ApiParamValidError(err)
	}

	storeUpdate := requestDto.StoreUpdate{}
	if err = ctx.Bind(&storeUpdate); err != nil {
		return errors.ApiParamValidError(err)
	}
	if err = storeUpdate.Validate(ctx); err != nil {
		return err
	}

	storeUpdate.No = storeNo

	err = service2.StoreService().Update(ctx.Request().Context(), storeUpdate)

	return ctx.NoContent(http.StatusOK)
}

func (StoreController) GetStores(ctx echo.Context) error {
	pageable := requestDto.GetPageableFromRequest(ctx)

	searchParams := requestDto.SearchStoreQueryParams{}

	result, err := service2.StoreService().GetStores(ctx.Request().Context(), searchParams, pageable)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (StoreController) Delete(ctx echo.Context) error {
	storeNo, err := strconv.ParseInt(ctx.Param("no"), 10, 64)

	if err != nil {
		return errors.ApiParamValidError(err)
	}

	err = service2.StoreService().Delete(ctx.Request().Context(), storeNo)

	return ctx.NoContent(http.StatusOK)
}
