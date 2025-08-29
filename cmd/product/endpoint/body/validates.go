package body

import (
	"fmt"
	"slices"

	"github.com/go-playground/validator/v10"
)

func (body *S_Body) validate() error {
	validate := validator.New()

	err := validate.Struct(*body)
	if err != nil {
		return err
	}

	if !slices.Contains(CONTENT_TYPE_ALLOWED, body.ContentType) {
		return fmt.Errorf("invalid endpoint._body.content_type %s | available options: %v", body.ContentType, CONTENT_TYPE_ALLOWED)
	}

	if body.ContentType != "application/x-www-form-urlencoded" && body.ContentType != "multipart/form-data" {
		if body.File == nil && body.Content == nil {
			return fmt.Errorf("invalid endpoint._body._file or endpoint._body._content | one of the two must be provided")
		}
	}

	if body.ContentType == "application/x-www-form-urlencoded" && body.FormFields == nil {
		return fmt.Errorf("invalid endpoint._body._form_fields | must be provided")
	}

	if body.ContentType == "multipart/form-data" && body.MultiPart == nil {
		return fmt.Errorf("invalid endpoint._body._multi_part | must be provided")
	}

	return nil
}
