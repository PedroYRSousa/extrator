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
	// TODO, ajustar
	"dirPath":              "must be dirPath",
	"filePath":             "must be filePath",
	"gt":                   "must be gt",
	"required_without_all": "must be required_without_all",
	"required_without":     "must be required_without",
	"excluded_without_all": "must be excluded_without_all",
	"excluded_without":     "must be excluded_without",
	"oneoffield":           "must be oneoffield",
	"contains":             "must be contains",
	"oneof":                "must be oneof",
	"url":                  "must be url",
}

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
