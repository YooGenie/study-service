package dto

import (
	"Menu/menu-service/common/errors"
	businessNumber "github.com/YooGenie/validateBusinessNumber"
	"github.com/labstack/echo"
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

if !businessNumber.ValidateBusinessNumber(a.BusinessRegistrationNumber) {
	err := errors.ErrNotValid
	return err
}
	return nil
}

//func validateBusinessRegistrationNumber(businessRegistrationNumber string) (err error) {
//
//	if len(businessRegistrationNumber) != 10 {
//		err = errors.ValidationError("사업자번호는 10자리입니다.")
//		return
//	}
//
//	return
//}

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

	if err := validateBusinessRegistrationNumber(a.BusinessRegistrationNumber); err != nil {
		return err
	}

	return nil
}

type SearchStoreQueryParams struct {
}
