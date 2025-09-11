package utils_structs

import (
	"extrator/pkg/env"
	"extrator/pkg/secret"
	"reflect"
)

func transformString(field reflect.Value) error {
	if field.Kind() == reflect.String {
		if field.CanSet() {
			if isEnv(field.String()) {
				value, err := env.Get(field.String())
				if err != nil {
					return err
				}

				field.SetString(value)
			} else if isSecret(field.String()) {
				value, err := secret.Get(field.String())
				if err != nil {
					return err
				}

				field.SetString(value)
			}
		}
	} else if field.Kind() == reflect.Map {
		for _, key := range field.MapKeys() {
			val := field.MapIndex(key)

			if val.Kind() == reflect.String {
				v := val.String()

				if isEnv(v) {
					value, err := env.Get(v)
					if err != nil {
						return err
					}

					field.SetMapIndex(key, reflect.ValueOf(value))
				} else if isSecret(v) {
					value, err := secret.Get(v)
					if err != nil {
						return err
					}

					field.SetMapIndex(key, reflect.ValueOf(value))
				}
			}
		}
	} else if field.Kind() == reflect.Struct {
		if field.CanAddr() {
			fieldAddr := field.Addr()
			err := Start(fieldAddr.Interface())
			if err != nil {
				return err
			}
		}
	} else if field.Kind() == reflect.Pointer {
		if !field.IsNil() {
			err := transformString(field.Elem())
			if err != nil {
				return err
			}
		}
	} else if field.Kind() == reflect.Slice {
		for i := 0; i < field.Len(); i++ {
			transformString(field.Index(i))
		}
	}

	return nil
}
