package dto

import (
	"github.com/labstack/echo"
	"menu-service/common/errors"
)

type StoreCreate struct {
	Id                         string `json:"id" validate:"required"`
	Password                   string `json:"password" validate:"required"`
	Mobile                     string `json:"mobile" validate:"required"`
	BusinessRegistrationNumber string `json:"businessRegistrationNumber" validate:"required"`
}

func (a StoreCreate) Validate(ctx echo.Context) error {

	//if err := ctx.Validate(a); err != nil {
	//	return err
	//}

	if err := validateBusinessRegistrationNumber(a.BusinessRegistrationNumber); err != nil {
		return err
	}

	return nil
}

func validateBusinessRegistrationNumber(businessRegistrationNumber string) (err error) {

	if len(businessRegistrationNumber) != 10 {
		err = errors.ValidationError("사업자번호는 10자리입니다.")
		return
	}

	return
}