package modules

import (
	"extrator/pkg/aws"
	"extrator/pkg/env"
	"reflect"
)

func transformString(field reflect.Value) error {
	if field.Kind() == reflect.String && field.CanSet() {
		if IsEnv(field.String()) {
			value, err := env.Get(field.String())
			if err != nil {
				return err
			}

			field.SetString(value)
		} else if IsSecret(field.String()) {
			value, err := aws.Get(field.String())
			if err != nil {
				return err
			}

			field.SetString(value)
		}
	} else if field.Kind() == reflect.Map && field.CanSet() {
		for _, key := range field.MapKeys() {
			val := field.MapIndex(key)

			if val.Kind() == reflect.String {
				v := val.String()

				if IsEnv(v) {
					value, err := env.Get(v)
					if err != nil {
						return err
					}
					field.SetMapIndex(key, reflect.ValueOf(value))
				} else if IsSecret(v) {
					value, err := aws.Get(v)
					if err != nil {
						return err
					}
					field.SetMapIndex(key, reflect.ValueOf(value))
				}
			}
		}
	} else if field.Kind() == reflect.Struct {
		fieldAddr := field
		if field.CanAddr() {
			fieldAddr = field.Addr()
			err := TransformString(fieldAddr.Interface())
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
	}

	return nil
}

func TransformString(structData any) error {
	val := reflect.ValueOf(structData).Elem()

	if val.Kind() != reflect.Struct {
		panic("TODO, alterar mensagem de erro")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		err := transformString(field)
		if err != nil {
			return err
		}
	}

	return nil
}
