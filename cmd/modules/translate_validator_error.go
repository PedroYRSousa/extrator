package modules

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func mapTranslate(tag string) string {
	switch tag {
	case "dir":
		return "The field must be a valid directory path"
	case "dirpath":
		return "The field must be a valid directory path"
	case "file":
		return "The field must be a valid file path"
	case "filepath":
		return "The field must be a valid file path"
	default:
		return tag
	}
}

func TranslateValidatorError(errs validator.ValidationErrors) error {
	for _, err := range errs {
		return fmt.Errorf("validation Error on field (%s): %s", err.Namespace(), mapTranslate(err.Tag()))
	}

	return nil
}
