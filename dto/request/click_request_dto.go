package dto

import (
	"github.com/labstack/echo/v4"
)

type ClickCreate struct {
	Click string `json:"click" validate:"required"`
}

func (a ClickCreate) Validate(ctx echo.Context) error {

	if err := ctx.Validate(a); err != nil {
		return err
	}

	return nil
}

type SearchClickQueryParams struct {
}
