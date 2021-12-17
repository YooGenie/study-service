package controllers

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"menu-service/common/errors"
	requestDto "menu-service/dto/request"
	"menu-service/menu/service"
	"strconv"

	"net/http"
)

type MenuController struct {
}

func (controller MenuController) Init(g *echo.Group) {
	g.POST("", controller.Create)
	g.GET("/:id", controller.GetMenuById)
	g.GET("", controller.GetMenu)
	g.PUT("/:id", controller.Update)
	g.DELETE("/:id", controller.Delete)
}

func (MenuController) Create(ctx echo.Context) error {
	var menuCreate = requestDto.MenuCreate{}

	if err := ctx.Bind(&menuCreate); err != nil {
		return errors.ApiParamValidError(err)
	}

	if err := menuCreate.Validate(ctx); err != nil {
		log.Errorf("Create Error:  %s", err.Error())
		return err
	}

	//서비스부분 연결
	err := service.MenuService().CreateMenu(ctx.Request().Context(), menuCreate)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (MenuController) GetMenuById(ctx echo.Context) error {
	menuId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.ApiParamValidError(err)
	}

	menu, _ := service.MenuService().GetMenuById(ctx.Request().Context(), menuId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, menu)

}

func (MenuController) GetMenu(ctx echo.Context) error {
	pageable := requestDto.GetPageableFromRequest(ctx)

	result, err := service.MenuService().GetMenu(ctx.Request().Context(), pageable)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}
func (MenuController) Update(ctx echo.Context) error {
	menuId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		return errors.ApiParamValidError(err)
	}

	menuUpdate := requestDto.MenuUpdate{}
	if err = ctx.Bind(&menuUpdate); err != nil {
		return errors.ApiParamValidError(err)
	}
	if err = menuUpdate.Validate(ctx); err != nil {
		return err
	}

	menuUpdate.Id = menuId

	err = service.MenuService().UpdateMenu(ctx.Request().Context(), menuUpdate)

	return ctx.JSON(http.StatusOK, nil)
}

func (MenuController) Delete(ctx echo.Context) error {
	menuId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.ApiParamValidError(err)
	}

	err = service.MenuService().DeleteMenu(ctx.Request().Context(), menuId)

	return ctx.JSON(http.StatusOK, nil)

}
