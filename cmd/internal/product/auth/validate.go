package auth

import (
	"errors"
	"extrator/internal/utils"
	"fmt"
	"slices"
)

func (ab *S_Basic) validate() error {
	ab.format()

	if ab.Username == "" {
		return errors.New("check auth.basic.username | value cannot be empty")
	}

	if !utils.CheckIsHardCodedSecretOrEnv(ab.Username) {
		return errors.New("check auth.basic.username | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	if ab.Password == "" {
		return errors.New("check auth.basic.password | value cannot be empty")
	}

	if !utils.CheckIsHardCodedSecretOrEnv(ab.Password) {
		return errors.New("check auth.basic.password | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	return nil
}

func (dd *S_DynamicDetails) validate() error {
	dd.format()

	if dd.Endpoint == "" {
		return errors.New("check auth.bearer.dynamic_details.endpoint | value cannot be empty")
	}

	if dd.Extract == "" {
		return errors.New("check auth.bearer.dynamic_details.extract | value cannot be empty")
	}

	err := dd.EndpointConfig.Validate()
	if err != nil {
		return fmt.Errorf("check auth.bearer.dynamic_details.endpoint_config | %v", err.Error())
	}

	return nil
}

func (ab *S_Bearer) validate() error {
	ab.format()

	if ab.Name == nil || *ab.Name == "" {
		return errors.New("check auth.bearer.name | value cannot be empty")
	}

	if ab.Prefix == nil || *ab.Prefix == "" {
		return errors.New("check auth.bearer.prefix | value cannot be empty")
	}

	if ab.Token == "" {
		return errors.New("check auth.bearer.token | value cannot be empty")
	}

	if !utils.CheckIsHardCodedSecretOrEnv(ab.Token) {
		return errors.New("check auth.bearer.token | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	err := ab.DynamicDetails.validate()
	if err != nil {
		return fmt.Errorf("check auth.bearer.dynamic_details | %v", err.Error())
	}

	return nil
}

func (apk *S_ApiKey) validate() error {
	apk.format()

	if apk.Name == "" {
		return errors.New("check auth.api_key.name | value cannot be empty")
	}

	if apk.Value == "" {
		return errors.New("check auth.api_key.value | value cannot be empty")
	}

	if !utils.CheckIsHardCodedSecretOrEnv(apk.Value) {
		return errors.New("check auth.api_key.value | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	if apk.Location == nil || !slices.Contains(AVAILABLE_LOCATIONS, *apk.Location) {
		return fmt.Errorf("check auth.api_key.location | available options: %v", AVAILABLE_LOCATIONS)
	}

	return nil
}

func (ab *S_Cookie) validate() error {
	ab.format()

	if ab.Endpoint == "" {
		return errors.New("check auth.cookie.endpoint | value cannot be empty")
	}

	if len(ab.Extract) == 0 {
		return errors.New("check auth.cookie.extract | value cannot be empty")
	}

	err := ab.EndpointConfig.Validate()
	if err != nil {
		return fmt.Errorf("check auth.cookie.endpoint_config | %v", err.Error())
	}

	return nil
}

func (a *S_Auth) Validate() error {
	a.format()

	if !slices.Contains(MODES_AVAILABLE, a.Mode) {
		return fmt.Errorf("check auth.mode | available options: %v", MODES_AVAILABLE)
	}

	switch a.Mode {
	case "basic":
		if a.Basic == nil {
			return errors.New("check auth.basic | basic cannot be null when mode is 'basic'")
		}
		return a.Basic.validate()
	case "bearer":
		if a.Bearer == nil {
			return errors.New("check auth.bearer | bearer cannot be null when mode is 'bearer'")
		}
		return a.Bearer.validate()
	case "api_key":
		if a.ApiKey == nil {
			return errors.New("check auth.api_key | api_key cannot be null when mode is 'api_key'")
		}
		return a.ApiKey.validate()
	case "cookie":
		if a.Cookie == nil {
			return errors.New("check auth.cookie | cookie cannot be null when mode is 'cookie'")
		}
		return a.Cookie.validate()
	}

	return nil
}
