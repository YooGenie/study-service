package controller

import (
	"net/http"
	"study-service/common/errors"
	requestDto "study-service/dto/request"
	"study-service/service"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type MemberController struct {
}

func (controller MemberController) Init(g *echo.Group) {
	g.POST("", controller.Create)
}

// Create
// @Tags 회원
// @Summary 멤버 등록
// @Description 함수에 대한 상세 내용 쓰기
// @Success 201 {object} nil
// @Router /api/member [post]
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
