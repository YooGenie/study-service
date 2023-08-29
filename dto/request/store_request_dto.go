package dto

import (
	"study-service/common/errors"

	"github.com/YooGenie/validate-business-number"
	"github.com/labstack/echo/v4"
)

type StoreCreate struct {
	Id                         string `json:"id" validate:"required"`
	Password                   string `json:"password" validate:"required"`
	Mobile                     string `json:"mobile" validate:"required"`
	BusinessRegistrationNumber string `json:"businessRegistrationNumber" validate:"required"`
}

func (a StoreCreate) Validate(ctx echo.Context) error {

	if err := ctx.Validate(a); err != nil {
		return err
	}

	if validate.BusinessNumber(a.BusinessRegistrationNumber) {
		err := errors.ErrNotValid
		return err
	}
	return nil
}

type StoreUpdate struct {
	No                         int64  `json:"no" validate:"required"`
	Id                         string `json:"id" validate:"required"`
	Password                   string `json:"password" validate:"required"`
	Mobile                     string `json:"mobile" validate:"required"`
	BusinessRegistrationNumber string `json:"businessRegistrationNumber" validate:"required"`
}

func (a StoreUpdate) Validate(ctx echo.Context) error {
	if err := ctx.Validate(a); err != nil {
		return err
	}

	if validate.BusinessNumber(a.BusinessRegistrationNumber) {
		err := errors.ErrNotValid
		return err
	}

	return nil
}

type SearchStoreQueryParams struct {
}
