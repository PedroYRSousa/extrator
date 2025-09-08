package validator

import (
	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	_ = v.RegisterValidation("required_without", requiredWithout, true)

	return v
}
