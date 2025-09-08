package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validação customizada
func requiredWithout(fl validator.FieldLevel) bool {
	count := 0
	parent := fl.Parent()
	params := strings.Split(fl.Param(), " ")

	if parent.Kind() == reflect.Ptr {
		if parent.IsNil() {
			return true
		}

		parent = parent.Elem()
	}

	if parent.FieldByName(fl.FieldName()).IsNil() {
		return true
	}

	for _, param := range params {
		if !parent.FieldByName(param).IsNil() {
			count += 1
		}
	}

	return count == 0
}
