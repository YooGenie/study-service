package dto

import (
	"github.com/labstack/echo"
)

type MemberCreate struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,gte=6,lte=100"`
	Mobile   string `json:"mobile" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Nickname string `json:"nickname"`
}

func (a MemberCreate) Validate(ctx echo.Context) (err error) {

	if err = ctx.Validate(a); err != nil {
		return
	}

	return
}



type SearchMemberQueryParams struct {
}
