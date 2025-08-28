package modules

import (
	"reflect"
	"strings"
)

func FormatString(structData any) {
	val := reflect.ValueOf(structData).Elem()

	if val.Kind() != reflect.Struct {
		panic("TODO, alterar mensagem de erro")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String && field.CanSet() {
			field.SetString(strings.TrimSpace(field.String()))
		}
	}
}
