package product

import (
	"github.com/go-playground/validator"
)

func (product *S_Product) validate() error {
	validate := validator.New()

	err := validate.Struct(*product)
	if err != nil {
		return err
	}

	return nil
}
