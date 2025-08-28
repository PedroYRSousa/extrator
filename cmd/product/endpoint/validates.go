package endpoint

import (
	"github.com/go-playground/validator"
)

func (endpoint *S_Endpoint) validate() error {
	validate := validator.New()

	err := validate.Struct(*endpoint)
	if err != nil {
		return err
	}

	return nil
}
