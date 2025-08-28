package auth

import (
	"errors"
	"extrator/internal/utils"
	"fmt"
	"slices"
)

func (basic *S_Basic) validate() error {
	if basic == nil {
		return errors.New("check auth.basic | basic cannot be null when mode is 'basic'")
	}

	basic.format()

	if basic.Username == "" {
		return errors.New("check auth.basic.username | value cannot be empty")
	}
	if !utils.CheckIsHardCodedSecretOrEnv(basic.Username) {
		return errors.New("check auth.basic.username | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	if basic.Password == "" {
		return errors.New("check auth.basic.password | value cannot be empty")
	}
	if !utils.CheckIsHardCodedSecretOrEnv(basic.Password) {
		return errors.New("check auth.basic.password | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	return nil
}

func (bearerDynamicDetails *S_BearerDynamicDetails) validate() error {
	if bearerDynamicDetails == nil {
		return errors.New("check auth.bearer.dynamic_details | dynamic_details cannot be null when mode is 'dynamic'")
	}

	bearerDynamicDetails.format()

	if bearerDynamicDetails.Endpoint == "" {
		return errors.New("check auth.bearer.dynamic_details.endpoint | value cannot be empty")
	}

	if bearerDynamicDetails.Extract == "" {
		return errors.New("check auth.bearer.dynamic_details.extract | value cannot be empty")
	}

	err := bearerDynamicDetails.EndpointConfig.Validate()
	if err != nil {
		return fmt.Errorf("check auth.bearer.dynamic_details.endpoint_config | %v", err.Error())
	}

	return nil
}

func (bearer *S_Bearer) validate() error {
	if bearer == nil {
		return errors.New("check auth.bearer | bearer cannot be null when mode is 'bearer'")
	}

	bearer.format()

	if bearer.Name == nil || *bearer.Name == "" {
		return errors.New("check auth.bearer.name | value cannot be empty")
	}

	if bearer.Prefix == nil || *bearer.Prefix == "" {
		return errors.New("check auth.bearer.prefix | value cannot be empty")
	}

	if bearer.Token == "" {
		return errors.New("check auth.bearer.token | value cannot be empty")
	}
	if !utils.CheckIsHardCodedSecretOrEnv(bearer.Token) {
		return errors.New("check auth.bearer.token | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	if *bearer.Dynamic {
		err := bearer.DynamicDetails.validate()
		if err != nil {
			return fmt.Errorf("check auth.bearer.dynamic_details | %v", err.Error())
		}
	}

	return nil
}

func (apiKey *S_ApiKey) validate() error {
	if apiKey == nil {
		return errors.New("check auth.api_key | api_key cannot be null when mode is 'api_key'")
	}

	apiKey.format()

	if apiKey.Name == "" {
		return errors.New("check auth.api_key.name | value cannot be empty")
	}

	if apiKey.Value == "" {
		return errors.New("check auth.api_key.value | value cannot be empty")
	}
	if !utils.CheckIsHardCodedSecretOrEnv(apiKey.Value) {
		return errors.New("check auth.api_key.value | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	if apiKey.Location == nil || !slices.Contains(AVAILABLE_LOCATIONS, *apiKey.Location) {
		return fmt.Errorf("invalid auth.api_key.location %q | available options: %v", apiKey.Location, AVAILABLE_LOCATIONS)
	}

	return nil
}

func (cookie *S_Cookie) validate() error {
	if cookie == nil {
		return errors.New("check auth.cookie | cookie cannot be null when mode is 'cookie'")
	}

	cookie.format()

	if cookie.Endpoint == "" {
		return errors.New("check auth.cookie.endpoint | value cannot be empty")
	}

	err := cookie.EndpointConfig.Validate()
	if err != nil {
		return fmt.Errorf("check auth.cookie.endpoint_config | %v", err.Error())
	}

	return nil
}

func (auth *S_Auth) Validate() error {
	var validators = map[authMode]func(*S_Auth) error{
		AUTH_MODE_BASIC:   func(_auth *S_Auth) error { return _auth.Basic.validate() },
		AUTH_MODE_BEARER:  func(_auth *S_Auth) error { return _auth.Bearer.validate() },
		AUTH_MODE_API_KEY: func(_auth *S_Auth) error { return _auth.ApiKey.validate() },
		AUTH_MODE_COOKIE:  func(_auth *S_Auth) error { return _auth.Cookie.validate() },
	}

	auth.format()

	validate, ok := validators[auth.Mode]
	if !ok {
		return fmt.Errorf("invalid auth.mode %q | available options: %v", auth.Mode, MODES_AVAILABLE)
	}

	return validate(auth)
}
