package modules

import (
	"reflect"
	"strings"
)

func formatString(field reflect.Value) {
	if field.Kind() == reflect.String && field.CanSet() {
		field.SetString(strings.TrimSpace(field.String()))
	} else if field.Kind() == reflect.Map && field.CanSet() {
		for _, key := range field.MapKeys() {
			val := field.MapIndex(key)

			if val.Kind() == reflect.String {
				field.SetMapIndex(key, reflect.ValueOf(strings.TrimSpace(val.String())))
			}
		}
	} else if field.Kind() == reflect.Struct {
		fieldAddr := field
		if field.CanAddr() {
			fieldAddr = field.Addr()
			FormatString(fieldAddr.Interface())
		}
	} else if field.Kind() == reflect.Pointer {
		if !field.IsNil() {
			formatString(field.Elem())
		}
	}
}

func FormatString(structData any) {
	val := reflect.ValueOf(structData).Elem()

	if val.Kind() != reflect.Struct {
		panic("TODO, alterar mensagem de erro")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		formatString(field)
	}
}
