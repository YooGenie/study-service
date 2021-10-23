package controllers

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"menu-service/common"
	"menu-service/dtos"
	"menu-service/services"
	"net/http"
	"strconv"
)

type MenuController struct {
}

func (controller MenuController) Init(g *echo.Group) {
	g.POST("", controller.Create)
	g.GET("/:Id", controller.GetMenuById)
	g.GET("", controller.GetMenu)
	g.PUT("/:Id", controller.Update)
	g.DELETE("/:Id", controller.Delete)
}

func (MenuController) Create(ctx echo.Context) error {
	var menuMake = dtos.MenuMake{}

	if err := ctx.Bind(&menuMake); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ApiError{
			Message: err.Error(),
		})
	}

	if err := ctx.Validate(&menuMake); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ApiError{
			Message: err.Error(),
		})
	}

	//서비스부분 연결
	Id, err := services.MenuService().CreateMenu(ctx.Request().Context(), menuMake)
	if err != nil {
		log.Errorf("Create Error:  %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, dtos.ApiError{
			Message: err.Error(),
		})
	}

	//응답하는 부분
	response := dtos.ResponseId{Id: Id}

	return ctx.JSON(http.StatusOK, response)
}

func (MenuController) GetMenuById(ctx echo.Context) error {
	menuId, _ := strconv.ParseInt(ctx.Param("Id"), 10, 64)

	menu, _ := services.MenuService().GetMenuById(ctx.Request().Context(), menuId)

	menuSummary := dtos.MenuSummary{
		Id:          menu.Id,
		Name:        menu.Name,
		Price:       menu.Price,
		CreatedBy:   menu.CreatedBy,
		CreatedAt:   menu.CreatedAt,
		UpdatedBy:   menu.UpdatedBy,
		UpdatedAt:   menu.UpdatedAt,
		Description: menu.Description,
	}

	return ctx.JSON(http.StatusOK, menuSummary)
}

func (MenuController) GetMenu(ctx echo.Context) error {
	pageable := dtos.GetPageableFromRequest(ctx)

	menus, totalCount, err := services.MenuService().GetMenu(ctx.Request().Context(), pageable)
	if err != nil {
		log.Errorf("GetMenu Error:  %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, dtos.ApiError{
			Message: err.Error(),
		})
	}
	menuSummaries := make([]dtos.MenuSummary, 0)

	for _, menu := range menus {
		menuSummaries = append(menuSummaries, dtos.MenuSummary{
			Id:          menu.Id,
			Name:        menu.Name,
			Price:       menu.Price,
			Description: menu.Description,
			CreatedBy:   menu.CreatedBy,
			CreatedAt:   menu.CreatedAt,
			UpdatedBy:   menu.UpdatedBy,
			UpdatedAt:   menu.UpdatedAt,
		})
	}

	resPageResult := dtos.PageResult{
		Result:     menuSummaries,
		TotalCount: totalCount,
	}

	return ctx.JSON(http.StatusOK, resPageResult)
}

func (MenuController) Update(ctx echo.Context) error {
	menuId, _ := strconv.ParseInt(ctx.Param("Id"), 10, 64)
	var menuMake dtos.MenuMake

	if err := ctx.Bind(&menuMake); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ApiError{
			Message: err.Error(),
		})
	}

	menuMake.Id = menuId

	menuId, _ = services.MenuService().UpdateMenu(ctx.Request().Context(), menuMake)

	return ctx.JSON(http.StatusOK, nil)
}

func (MenuController) Delete(ctx echo.Context) error {
	menuId, err := strconv.ParseInt(ctx.Param("Id"), 10, 64)
	if err != nil {
		log.Errorf("Delete Error: %s", err.Error())

		return ctx.JSON(http.StatusBadRequest, dtos.ApiError{
			Message: err.Error(),
		})
	}

	err = services.MenuService().DeleteMenu(ctx.Request().Context(), menuId)
	if err != nil {
		log.Errorf("Delete Error:  %s", err.Error())
		if err == common.ErrAuthorization {
			return ctx.JSON(http.StatusInternalServerError, common.APIErrorAuthorization)
		}
		return ctx.JSON(http.StatusInternalServerError, dtos.ApiError{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, nil)
}
