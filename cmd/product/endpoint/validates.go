package endpoint

import (
	"fmt"
	"slices"

	"github.com/go-playground/validator/v10"
)

func (endpoint *S_Endpoint) validate() error {
	validate := validator.New()

	err := validate.Struct(*endpoint)
	if err != nil {
		return err
	}

	if !slices.Contains(METHODS_ALLOWED, *endpoint.Method) {
		return fmt.Errorf("invalid endpoint._method %q | available options: %v", *endpoint.Method, METHODS_ALLOWED)
	}

	if !slices.Contains(RESPONSE_FORMAT_ALLOWED, *endpoint.ResponseFormat) {
		return fmt.Errorf("invalid endpoint._response_format %q | available options: %v", *endpoint.ResponseFormat, RESPONSE_FORMAT_ALLOWED)
	}

	return nil
}
