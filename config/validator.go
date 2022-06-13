package config

import (
	"github.com/go-playground/validator/v10"
	val "study-service/validator"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) (err error) {
	return cv.validator.Struct(i)
}

func RegisterValidator() *CustomValidator {
	customValidator := validator.New()
	customValidator.RegisterValidation("date12", val.ValidateDate12)

	return &CustomValidator{validator: customValidator}
}
