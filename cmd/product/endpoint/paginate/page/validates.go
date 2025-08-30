package page

import (
	"fmt"
	"slices"

	"github.com/go-playground/validator/v10"
)

func (page *S_PaginatePage) validate() error {
	validate := validator.New()

	err := validate.Struct(*page)
	if err != nil {
		return err
	}

	if !slices.Contains(DIRECTION_ALLOWED, *page.Direction) {
		return fmt.Errorf("invalid endpoint._paginate.mode %s | available options: %v", *page.Direction, DIRECTION_ALLOWED)
	}

	if !slices.Contains(LOCATION_ALLOWED, *page.Location) {
		return fmt.Errorf("invalid endpoint._paginate.mode %s | available options: %v", *page.Direction, LOCATION_ALLOWED)
	}

	return nil
}
