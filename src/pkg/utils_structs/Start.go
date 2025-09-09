package utilsstructs

import "reflect"

func Start(structData any) error {
	val := reflect.ValueOf(structData).Elem()

	if val.Kind() == reflect.Pointer || val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}

		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		panic("ERROR Interno, verificar struct_inspector.Start")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		formatString(field)
		err := transformString(field)
		if err != nil {
			return err
		}
	}

	return nil
}
