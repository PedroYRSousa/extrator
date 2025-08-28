package modules

import (
	"extrator/pkg/aws"
	"extrator/pkg/env"
	"reflect"
)

func TransformString(structData any) error {
	val := reflect.ValueOf(structData).Elem()

	if val.Kind() != reflect.Struct {
		panic("TODO, alterar mensagem de erro")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
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
		}
	}

	return nil
}
