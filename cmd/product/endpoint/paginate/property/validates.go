package property

import (
	"github.com/go-playground/validator/v10"
)

func (paginateProperty *S_PaginateProperty) validate() error {
	validate := validator.New()

	err := validate.Struct(*paginateProperty)
	if err != nil {
		return err
	}

	return nil
}
