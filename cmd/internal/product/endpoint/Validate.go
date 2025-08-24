package endpoint

import (
	"errors"
	"fmt"
	"net/url"

	"golang.org/x/exp/slices"
)

func (e *S_Endpoint) Validate() error {
	e.format()

	if e.URI == "" {
		return errors.New("check endpoint.uri | uri cannot be empty")
	}

	uri, err := url.Parse(e.URI)
	if err != nil {
		return fmt.Errorf("check endpoint.uri | error: %v", err)
	}

	// Não pode ter query params na URI
	if uri.Query().Encode() != "" {
		return errors.New("check endpoint.uri | query parameters must be defined in endpoint_config.query_params")
	}

	if !slices.Contains(RESPONSES_FORMAT_AVAILABLE, e.ResponseFormat) {
		return fmt.Errorf("check endpoint.response_format | available options: %v", RESPONSES_FORMAT_AVAILABLE)
	}

	// LimitExtract, pode ser qualquer valor

	err = e.EndpointConfig.Validate()
	if err != nil {
		return err
	}

	return nil
}
