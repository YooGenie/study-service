package dto

import (
	"github.com/labstack/echo"
)

type AdminSignIn struct {
	Id       string `json:"id"`
	Password string `json:"password" validate:"gte=6,lte=100"`
}

func (r AdminSignIn) Validate(ctx echo.Context) (err error) {
	//if err = ctx.Validate(r); err != nil {
	//	return
	//}
	return
}
