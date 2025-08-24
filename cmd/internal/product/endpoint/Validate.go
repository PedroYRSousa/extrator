package endpoint

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/exp/slices"
)

func (e *S_Endpoint) Validate() error {
	e.format()

	if e.URI == "" {
		return errors.New("check endpoint.uri | value cannot be empty")
	}

	uri, err := url.Parse(e.URI)
	if err != nil {
		return fmt.Errorf("check endpoint.uri | error: %v", err)
	}

	// Não pode ter query params na URI
	if uri.Query().Encode() != "" {
		return errors.New("check endpoint.uri | query parameters must be defined in endpoint_config.query_params")
	}

	// Foco na extração de dados
	methodsAvailable := []string{http.MethodGet, http.MethodPost}
	if !slices.Contains(methodsAvailable, e.Method) {
		return fmt.Errorf("check endpoint.method | available options: %v", methodsAvailable)
	}

	// Por hora somente esses formatos são suportados para extração de dados
	if !slices.Contains(responsesFormatAvailable, e.ResponseFormat) {
		return fmt.Errorf("check endpoint.response_format | available options: %v", responsesFormatAvailable)
	}

	err = e.EndpointConfig.Validate()
	if err != nil {
		return err
	}

	return nil
}
