package endpoint

import (
	"errors"
	"extrator/internal/utils"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/exp/slices"
)

func (e *S_Endpoint) Validate() error {
	if e.URI == "" {
		return errors.New("invalid endpoint uri | Check endpoint.uri | value cannot be empty")
	}

	uri, err := url.Parse(e.URI)
	if err != nil {
		return fmt.Errorf("invalid endpoint uri | Check endpoint.uri | error: %v", err)
	}

	// Não pode ter query params na URI
	if uri.Query().Encode() != "" {
		return errors.New("invalid endpoint uri | Check endpoint.uri | query parameters must be defined in endpoint.query_params")
	}

	// Foco na extração de dados
	methodsAvailable := []string{http.MethodGet, http.MethodPost}
	if !slices.Contains(methodsAvailable, e.Method) {
		return fmt.Errorf("invalid endpoint method | Check endpoint.method | available options: %v", methodsAvailable)
	}

	// Por hora somente esses formatos são suportados para extração de dados
	if !slices.Contains(responsesFormatAvailable, e.ResponseFormat) {
		return fmt.Errorf("invalid endpoint response format | Check endpoint.response_format | available options: %v", responsesFormatAvailable)
	}

	err = e.EndpointConfig.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (r *S_Retry) validate() error {
	if r.Attempts < 0 {
		return errors.New("invalid endpoint config retry attempts | Check endpoint.endpoint_config.retry.attempts | value must be greater than or equal to 0")
	}

	if r.DelayInSeconds < 0 {
		return errors.New("invalid endpoint config retry delay in seconds | Check endpoint.endpoint_config.retry.delay_in_seconds | value must be greater than or equal to 0")
	}

	if *r.BackoffFactor < 0 {
		return errors.New("invalid endpoint config retry backoff factor | Check endpoint.endpoint_config.retry.backoff_factor | value must be greater than or equal to 0")
	}

	return nil
}

func (ec *S_EndpointConfig) Validate() error {
	if ec.WaitingTimeErrorInSeconds < 0 {
		return errors.New("invalid endpoint config waiting time in error seconds | Check endpoint.endpoint_config.waiting_time_in_error_in_seconds | value must be greater than or equal to 0")
	}

	if ec.WaitingTimeInSeconds < 0 {
		return errors.New("invalid endpoint config waiting time in seconds | Check endpoint.endpoint_config.waiting_time_in_seconds | value must be greater than or equal to 0")
	}

	if ec.Retry != nil {
		err := ec.Retry.validate()
		if err != nil {
			return err
		}
	}

	if ec.Headers != nil {
		for key, value := range *(ec.Headers) {
			if strings.TrimSpace(key) == "" {
				return errors.New("invalid endpoint headers | Check endpoint.headers | header names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("invalid endpoint body | Check endpoint.body | body parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
				}
			}
		}

	}

	if ec.Body != nil {
		for key, value := range *(ec.Body) {
			if strings.TrimSpace(key) == "" {
				return errors.New("invalid endpoint body | Check endpoint.body | body parameter names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("invalid endpoint body | Check endpoint.body | body parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
				}
			}
		}
	}

	if ec.TimeoutInSeconds != nil {
		ec.TimeoutInSeconds = new(int)
		*ec.TimeoutInSeconds = -1
	}

	// if ec.QueryParams != nil {
	// 	for _, queryParam := range *(ec.QueryParams) {
	// 		err := queryParam.validate()
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	return nil
}
