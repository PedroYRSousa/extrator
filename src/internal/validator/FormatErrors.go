package validator

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var mappingTag = map[string]string{
	"required":   "is required",
	"printascii": "must contain only printable ASCII characters",
	"uppercase":  "must be uppercase",
	"lowercase":  "must be lowercase",
}

// FormatErrors returns short English messages like:
// "ABC (yaml: path_products.teste.abc) is required"
func formatErros(err error, s any) []string {
	var errs []string
	if err == nil {
		return errs
	}

	typ := reflect.TypeOf(s)
	if typ == nil {
		errs = append(errs, err.Error())
		return errs
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range validationErrs {
			yamlPath := buildYamlPath(fe, typ)
			msg := fmt.Sprintf("%s | %s", yamlPath, mappingTag[fe.Tag()])
			errs = append(errs, msg)
		}
	} else {
		errs = append(errs, err.Error())
	}

	return errs
}
func FormatErrors(err error, s interface{}) string {
	errString := "ERRORS:\n"
	errs := formatErros(err, s)

	for _index, _err := range errs {
		if _index != len(errs)-1 {
			errString += fmt.Sprintf("\t%s\n", _err)
		} else {
			errString += fmt.Sprintf("\t%s", _err)
		}
	}

	return errString
}
