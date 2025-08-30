package paginate

import (
	"fmt"
	"slices"

	"github.com/go-playground/validator/v10"
)

func (paginate *S_Paginate) validate() error {
	validate := validator.New()

	err := validate.Struct(*paginate)
	if err != nil {
		return err
	}

	if !slices.Contains(MODE_ALLOWED, paginate.Mode) {
		return fmt.Errorf("invalid endpoint._paginate.mode %s | available options: %v", paginate.Mode, MODE_ALLOWED)
	}

	return nil
}
