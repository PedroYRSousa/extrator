package link_header

import (
	"github.com/go-playground/validator/v10"
)

func (paginateLinkHeader *S_PaginateLinkHeader) validate() error {
	validate := validator.New()

	err := validate.Struct(*paginateLinkHeader)
	if err != nil {
		return err
	}

	return nil
}
