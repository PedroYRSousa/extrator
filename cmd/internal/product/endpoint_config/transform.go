package endpoint_config

import (
	"extrator/internal/utils"
	"extrator/pkg/env"
)

func (r *S_Retry) transform() error {
	return nil
}

func (ec *S_EndpointConfig) Transform() error {
	if ec.Headers != nil {
		for k, v := range *ec.Headers {
			if utils.IsEnv(v) {
				newValue, err := env.Get(v)
				if err != nil {
					return err
				}
				(*ec.Headers)[k] = newValue
			} else if utils.IsSecret(v) {
				// TODO, Pegar da aws secretmanager
			}
		}
	}

	if ec.Body != nil {
		for k, v := range *ec.Body {
			if utils.IsEnv(v) {
				newValue, err := env.Get(v)
				if err != nil {
					return err
				}
				(*ec.Body)[k] = newValue
			} else if utils.IsSecret(v) {
				// TODO, Pegar da aws secretmanager
			}
		}
	}

	if ec.QueryParams != nil {
		for k, v := range *ec.QueryParams {
			if utils.IsEnv(v) {
				newValue, err := env.Get(v)
				if err != nil {
					return err
				}
				(*ec.QueryParams)[k] = newValue
			} else if utils.IsSecret(v) {
				// TODO, Pegar da aws secretmanager
			}
		}
	}

	err := ec.Retry.transform()
	if err != nil {
		return err
	}

	return nil
}
