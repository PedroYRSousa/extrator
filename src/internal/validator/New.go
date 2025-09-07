package validator

import (
	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	v := validator.New()

	// exemplo de validação customizada
	_ = v.RegisterValidation("tag-here", func(fl validator.FieldLevel) bool {
		return fl.Field().String() != "forbidden"
	})

	return v
}
