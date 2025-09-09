package utilsstructs

import (
	"reflect"
	"strings"
)

func formatString(field reflect.Value) {
	if field.Kind() == reflect.String {
		if field.CanSet() {
			field.SetString(strings.TrimSpace(field.String()))
		}
	} else if field.Kind() == reflect.Map {
		for _, key := range field.MapKeys() {
			val := field.MapIndex(key)

			if val.Kind() == reflect.String {
				field.SetMapIndex(key, reflect.ValueOf(strings.TrimSpace(val.String())))
			}
		}
	} else if field.Kind() == reflect.Struct {
		if field.CanAddr() {
			fieldAddr := field.Addr()
			Start(fieldAddr.Interface())
		}
	} else if field.Kind() == reflect.Pointer {
		if !field.IsNil() {
			formatString(field.Elem())
		}
	} else if field.Kind() == reflect.Slice {
		for i := 0; i < field.Len(); i++ {
			formatString(field.Index(i))
		}
	}
}
