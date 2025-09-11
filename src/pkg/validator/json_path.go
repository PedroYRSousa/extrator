package validator

import (
	"log"
	"reflect"

	"github.com/PaesslerAG/jsonpath"
	"github.com/go-playground/validator/v10"
)

func jsonPath(fl validator.FieldLevel) bool {
	parent := fl.Parent()

	if parent.Kind() == reflect.Ptr {
		if parent.IsNil() {
			return true
		}

		parent = parent.Elem()
	}

	if parent.FieldByName(fl.FieldName()).IsNil() {
		return false
	}

	_, err := jsonpath.New(parent.FieldByName(fl.FieldName()).String())
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
