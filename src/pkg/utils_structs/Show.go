package utils_structs

import (
	"fmt"
	"reflect"
	"strings"
)

func getValuePointer(v reflect.Value) *reflect.Value {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	return &v
}

func showStruct(v *reflect.Value, structFieldType *reflect.StructField, deep uint) {
	if v == nil {
		return
	}

	name := v.Type().Name()
	if structFieldType != nil {
		if tag := structFieldType.Tag.Get("yaml"); tag != "" {
			name = tag
		} else {
			name = structFieldType.Name
		}
	}

	structName := strings.ReplaceAll(name, "S_", "")
	fmt.Printf("%s%s {\n", strings.Repeat("  ", int(deep)), structName)
	for i := 0; i < v.NumField(); i++ {
		fieldVal := getValuePointer(v.Field(i))
		fieldType := v.Type().Field(i)
		if fieldVal == nil {
			continue
		}

		switch fieldVal.Kind() {
		case reflect.Struct:
			showStruct(fieldVal, &fieldType, deep+1)
		case reflect.Slice:
			showSlice(fieldVal, &fieldType, deep+1)
		case reflect.Map:
			showMap(fieldVal, &fieldType, deep+1)
		default:
			showGeneral(fieldVal, &fieldType, deep+1)
		}
	}
	fmt.Printf("%s}\n", strings.Repeat("  ", int(deep)))
}

func showSlice(v *reflect.Value, structFieldType *reflect.StructField, deep uint) {
	if v == nil {
		return
	}

	name := structFieldType.Name
	if tag := structFieldType.Tag.Get("yaml"); tag != "" {
		name = tag
	}

	fmt.Printf("%s%s [\n", strings.Repeat("  ", int(deep)), name)
	for i := 0; i < v.Len(); i++ {
		elem := getValuePointer(v.Index(i))
		if elem == nil {
			continue
		}

		switch elem.Kind() {
		case reflect.Struct:
			showStruct(elem, nil, deep+1)
		case reflect.Slice:
			showSlice(elem, structFieldType, deep+1)
		case reflect.Map:
			showMap(elem, structFieldType, deep+1)
		default:
			showGeneral(elem, structFieldType, deep+1)
		}
	}
	fmt.Printf("%s]\n", strings.Repeat("  ", int(deep)))
}

func showMap(v *reflect.Value, structFieldType *reflect.StructField, deep uint) {
	if v == nil {
		return
	}

	for _, key := range v.MapKeys() {
		val := getValuePointer(v.MapIndex(key))
		if val == nil {
			continue
		}

		fmt.Printf("%s%v: ", strings.Repeat("  ", int(deep+1)), key.Interface())

		switch val.Kind() {
		case reflect.Struct:
			fmt.Println()
			showStruct(val, nil, deep+1)
		case reflect.Slice:
			fmt.Println()
			showSlice(val, structFieldType, deep+1)
		case reflect.Map:
			fmt.Println()
			showMap(val, structFieldType, deep+1)
		default:
			fmt.Printf("\"%v\"\n", val.Interface())
		}
	}
}

func showGeneral(v *reflect.Value, structFieldType *reflect.StructField, deep uint) {
	if v == nil {
		return
	}

	name := structFieldType.Name
	if tag := structFieldType.Tag.Get("yaml"); tag != "" {
		name = tag
	}

	fmt.Printf("%s%s: \"%v\"\n", strings.Repeat("  ", int(deep)), name, v.Interface())
}

func Show(target any, deep uint) {
	v := getValuePointer(reflect.ValueOf(target))
	if v == nil || !v.IsValid() {
		return
	}

	if v.Kind() == reflect.Struct {
		showStruct(v, nil, deep)
	}
}
