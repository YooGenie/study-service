package dto

import (
	"github.com/labstack/echo"
	"menu-service/common/errors"
	"time"
)

//요청DTO
type MenuCreate struct {
	Name        string `json:"name" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Description string `json:"description"`
	//TestNumeric string `json:"testNumeric" validate:"numeric"`
}

func (a MenuCreate) Validate(ctx echo.Context) error {

	//if err := ctx.Validate(a); err != nil {
	//	return err
	//}

	if err := validatePrice(a.Price); err != nil {
		return err
	}

	return nil
}

func validatePrice(price int64) (err error) {

	if price == int64(0) || price < int64(0) {
		err = errors.ValidationError("가격을 입력해주세오")
		return
	}

	return
}

type MenuSummary struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedBy   string    `json:"updatedBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
