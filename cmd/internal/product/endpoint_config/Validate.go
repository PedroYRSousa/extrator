package endpointconfig

import (
	"errors"
	"extrator/internal/utils"
	"strings"
)

func (r *S_Retry) validate() error {
	r.format()

	if *r.Attempts < 0 {
		return errors.New("check endpoint_config.retry.attempts | value must be greater than or equal to 0")
	}

	if *r.DelayInSeconds < 0 {
		return errors.New("check endpoint_config.retry.delay_in_seconds | value must be greater than or equal to 0")
	}

	if r.BackoffFactor != nil && *r.BackoffFactor < 0 {
		return errors.New("check endpoint_config.retry.backoff_factor | value must be greater than or equal to 0")
	}

	return nil
}

func (ec *S_EndpointConfig) Validate() error {
	ec.format()

	if *ec.WaitingTimeErrorInSeconds < 0 {
		return errors.New("invalid endpoint config waiting time in error seconds | Check endpoint_config.waiting_time_in_error_in_seconds | value must be greater than or equal to 0")
	}

	if *ec.WaitingTimeInSeconds < 0 {
		return errors.New("invalid endpoint config waiting time in seconds | Check endpoint_config.waiting_time_in_seconds | value must be greater than or equal to 0")
	}

	err := ec.Retry.validate()
	if err != nil {
		return err
	}

	if ec.Headers != nil {
		for key, value := range *(ec.Headers) {
			if strings.TrimSpace(key) == "" {
				return errors.New("invalid endpoint headers | Check endpoint_config.headers | header names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("invalid endpoint body | Check endpoint_config.headers | body parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
				}
			}
		}
	}

	if ec.Body != nil {
		for key, value := range *(ec.Body) {
			if strings.TrimSpace(key) == "" {
				return errors.New("invalid endpoint body | Check endpoint_config.body | body parameter names cannot be empty")
			}

			if value != "" {
				if !utils.CheckIsHardCodedSecretOrEnv(value) {
					return errors.New("invalid endpoint body | Check endpoint_config.body | body parameter values must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
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
