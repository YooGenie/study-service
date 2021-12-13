package controllers

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"menu-service/common/errors"
	requestDto "menu-service/dto/request"
	"menu-service/menu/service"

	"net/http"
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
	var menuCreate= requestDto.MenuCreate{}

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
	//menuId, _ := strconv.ParseInt(ctx.Param("Id"), 10, 64)
	//
	//menu, _ := services.MenuService().GetMenuById(ctx.Request().Context(), menuId)
	//
	//menuSummary := dto2.MenuSummary{
	//	Id:          menu.Id,
	//	Name:        menu.Name,
	//	Price:       menu.Price,
	//	CreatedBy:   menu.CreatedBy,
	//	CreatedAt:   menu.CreatedAt,
	//	UpdatedBy:   menu.UpdatedBy,
	//	UpdatedAt:   menu.UpdatedAt,
	//	Description: menu.Description,
	//}

	//return ctx.JSON(http.StatusOK, menuSummary)
	return nil
}

func (MenuController) GetMenu(ctx echo.Context) error {
	//pageable := response.GetPageableFromRequest(ctx)
	//
	//menus, totalCount, err := services.MenuService().GetMenu(ctx.Request().Context(), pageable)
	//if err != nil {
	//	log.Errorf("GetMenu Error:  %s", err.Error())
	//	return ctx.JSON(http.StatusInternalServerError, response.ApiError{
	//		Message: err.Error(),
	//	})
	//}
	//menuSummaries := make([]dto2.MenuSummary, 0)
	//
	//for _, menu := range menus {
	//	menuSummaries = append(menuSummaries, dto2.MenuSummary{
	//		Id:          menu.Id,
	//		Name:        menu.Name,
	//		Price:       menu.Price,
	//		Description: menu.Description,
	//		CreatedBy:   menu.CreatedBy,
	//		CreatedAt:   menu.CreatedAt,
	//		UpdatedBy:   menu.UpdatedBy,
	//		UpdatedAt:   menu.UpdatedAt,
	//	})
	//}
	//
	//resPageResult := response.PageResult{
	//	Result:     menuSummaries,
	//	TotalCount: totalCount,
	//}
	//
	//return ctx.JSON(http.StatusOK, resPageResult)
	return nil
}

func (MenuController) Update(ctx echo.Context) error {
	//menuId, _ := strconv.ParseInt(ctx.Param("Id"), 10, 64)
	//var menuMake dto2.MenuMake
	//
	//if err := ctx.Bind(&menuMake); err != nil {
	//	return ctx.JSON(http.StatusBadRequest, response.ApiError{
	//		Message: err.Error(),
	//	})
	//}
	//
	//menuMake.Id = menuId
	//
	//menuId, _ = services.MenuService().UpdateMenu(ctx.Request().Context(), menuMake)
	//
	//return ctx.JSON(http.StatusOK, nil)

	return nil
}

func (MenuController) Delete(ctx echo.Context) error {
	//menuId, err := strconv.ParseInt(ctx.Param("Id"), 10, 64)
	//if err != nil {
	//	log.Errorf("Delete Error: %s", err.Error())
	//
	//	return ctx.JSON(http.StatusBadRequest, response.ApiError{
	//		Message: err.Error(),
	//	})
	//}
	//
	//err = services.MenuService().DeleteMenu(ctx.Request().Context(), menuId)
	//if err != nil {
	//	log.Errorf("Delete Error:  %s", err.Error())
	//	if err == common.ErrAuthorization {
	//		return ctx.JSON(http.StatusInternalServerError, common.ErrAuthorization)
	//	}
	//	return ctx.JSON(http.StatusInternalServerError, response.ApiError{
	//		Message: err.Error(),
	//	})
	//}
	//
	//return ctx.JSON(http.StatusOK, nil)

	return nil
}
