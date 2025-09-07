package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// buildYamlPath builds a dotted yaml path for the field error.
// It prefixes with the root's PathProducts yaml tag when present,
// producing e.g. "path_products.teste.abc".
func buildYamlPath(fe validator.FieldError, rootType reflect.Type) string {
	parts := strings.Split(fe.StructNamespace(), ".") // e.g. ["S_Configs","Teste","ABC"]
	if len(parts) <= 1 {
		return strings.ToLower(fe.Field())
	}

	yamlParts := make([]string, 0, len(parts)-1)
	t := rootType
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for _, part := range parts[1:] { // skip root type name
		f, ok := t.FieldByName(part)
		var yamlTag string
		if ok {
			yamlTag = f.Tag.Get("yaml")
			if yamlTag == "" {
				yamlTag = strings.ToLower(f.Name)
			}
		} else {
			// fallback: use lowercased part name
			yamlTag = strings.ToLower(part)
		}

		yamlParts = append(yamlParts, yamlTag)

		// descend into nested struct types when possible
		if ok {
			ft := f.Type
			if ft.Kind() == reflect.Ptr {
				ft = ft.Elem()
			}
			if ft.Kind() == reflect.Struct {
				t = ft
			}
		}
	}

	// prefix with root PathProducts yaml tag if present
	if prefix := findRootYamlPrefix(rootType); prefix != "" {
		if len(yamlParts) == 0 || yamlParts[0] != prefix {
			yamlParts = append([]string{prefix}, yamlParts...)
		}
	}

	return strings.Join(yamlParts, ".")
}

// findRootYamlPrefix looks for a field named PathProducts (preferred)
// or a field with yaml:"path_products" and returns its yaml tag.
// Returns empty string when not found.
func findRootYamlPrefix(t reflect.Type) string {
	if t == nil {
		return ""
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// prefer exact field name PathProducts
	if f, ok := t.FieldByName("PathProducts"); ok {
		if tag := f.Tag.Get("yaml"); tag != "" {
			return tag
		}
	}

	// fallback: look for a field whose yaml tag equals "path_products"
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("yaml") == "path_products" {
			return "path_products"
		}
	}

	return ""
}
