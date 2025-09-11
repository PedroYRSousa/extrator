package validator

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func New() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	err := v.RegisterValidation("required_without", requiredWithout, true)
	if err != nil {
		log.Fatalln("Internal error | validator.New", err)
	}

	err = v.RegisterValidation("json_path", jsonPath, true)
	if err != nil {
		log.Fatalln("Internal error | validator.New", err)
	}

	return v
}
