package endpoint_config

import (
	"errors"
	"extrator/internal/utils"
	"fmt"
	"slices"
	"strings"
)

func (r *S_Retry) validate() error {
	r.format()

	return nil
}

func (ec *S_EndpointConfig) Validate() error {
	ec.format()

	if !slices.Contains(METHODS_AVAILABLE, *ec.Method) {
		return fmt.Errorf("check endpoint_config.method | available options: %v", METHODS_AVAILABLE)
	}

	err := ec.Retry.validate()
	if err != nil {
		return err
	}

	if ec.Headers != nil {
		for key, value := range *(ec.Headers) {
			if strings.TrimSpace(key) == "" {
				return errors.New("check endpoint_config.headers | header names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("check endpoint_config.headers | body parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
				}
			}
		}
	}

	if ec.Body != nil {
		for key, value := range *(ec.Body) {
			if strings.TrimSpace(key) == "" {
				return errors.New("check endpoint_config.body | body parameter names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("check endpoint_config.body | body parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
				}
			}
		}
	}

	if ec.QueryParams != nil {
		for key, value := range *(ec.QueryParams) {
			if strings.TrimSpace(key) == "" {
				return errors.New("check endpoint_config.query_params | query_params parameter names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("check endpoint_config.query_params | query_params parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
				}
			}
		}
	}

	return nil
}
