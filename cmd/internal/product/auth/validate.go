package auth

import (
	"errors"
	"extrator/internal/utils"
	"fmt"
	"net/http"
	"slices"
)

func (a *S_Auth) Validate() error {
	modesAvailable := []string{"none", "basic", "bearer", "bearer_dynamic", "cookie", "api_key"}
	if !slices.Contains(modesAvailable, a.AuthMode) {
		return fmt.Errorf("invalid auth mode | Check auth.auth_mode | available options: %v", modesAvailable)
	}

	switch a.AuthMode {
	case "basic":
		if a.AuthBasic == nil {
			return errors.New("invalid auth basic | Check auth.auth_basic | auth_basic cannot be null when auth_mode is 'basic'")
		}
		return a.AuthBasic.validate()
	case "bearer":
		if a.AuthBearer == nil {
			return errors.New("invalid auth bearer | Check auth.auth_bearer | auth_bearer cannot be null when auth_mode is 'bearer'")
		}
		return a.AuthBearer.validate()
	case "api_key":
		if a.AuthApiKey == nil {
			return errors.New("invalid auth api key | Check auth.auth_api_key | auth_api_key cannot be null when auth_mode is 'api_key'")
		}
		return a.AuthApiKey.validate()
	case "bearer_dynamic":
		if a.AuthBearerDynamic == nil {
			return errors.New("invalid auth bearer dynamic | Check auth.auth_bearer_dynamic | auth_bearer_dynamic cannot be null when auth_mode is 'bearer dynamic'")
		}
		return a.AuthBearerDynamic.validate()
	case "cookie":
		if a.AuthCookie == nil {
			return errors.New("invalid auth cookie | Check auth.auth_cookie | auth_cookie cannot be null when auth_mode is 'cookie'")
		}
		return a.AuthCookie.validate()
	}

	return nil
}

func (ab *S_AuthBasic) validate() error {
	if ab.Username == "" {
		return errors.New("invalid auth basic username | Check auth.auth_basic.username | value cannot be empty")
	}

	if !utils.CheckIsHardCodedSecretOrEnv(ab.Username) {
		return errors.New("invalid auth basic username | Check auth.auth_basic.username | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	if ab.Password == "" {
		return errors.New("invalid auth basic password | Check auth.auth_basic.password | value cannot be empty")
	}

	if !utils.CheckIsHardCodedSecretOrEnv(ab.Password) {
		return errors.New("invalid auth basic password | Check auth.auth_basic.password | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	return nil
}

func (ab *S_AuthBearer) validate() error {
	if ab.Name == "" {
		ab.Name = "Authorization"
	}

	if ab.Prefix == "" {
		ab.Prefix = "Bearer"
	}

	if ab.Token == "" {
		return errors.New("invalid auth bearer token | Check auth.auth_bearer.token | value cannot be empty")
	}

	if !utils.CheckIsHardCodedSecretOrEnv(ab.Token) {
		return errors.New("invalid auth bearer token | Check auth.auth_bearer.token | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	return nil
}

func (apk *S_AuthApiKey) validate() error {
	if apk.Name == "" {
		return errors.New("invalid auth api key name | Check auth.auth_api_key.name | value cannot be empty")
	}

	availableLocations := []string{"header", "query_param"}
	if !slices.Contains(availableLocations, apk.Location) {
		return fmt.Errorf("invalid auth api key location | Check auth.auth_api_key.location | available options: %v", availableLocations)
	}

	if apk.Value == "" {
		return errors.New("invalid auth bearer token | Check auth.auth_bearer.token | value cannot be empty")
	}

	if !utils.CheckIsHardCodedSecretOrEnv(apk.Value) {
		return errors.New("invalid auth bearer token | Check auth.auth_bearer.token | value must be a hardcoded value, secret(SECRET_NAME) or env(ENV_VAR)")
	}

	return nil
}

func (abd *S_AuthBearerDynamic) validate() error {
	if abd.Name == "" {
		abd.Name = "Authorization"
	}

	if abd.Prefix == "" {
		abd.Prefix = "Bearer"
	}

	if abd.Endpoint == "" {
		return errors.New("invalid auth bearer dynamic endpoint | Check auth.auth_bearer_dynamic.endpoint | value cannot be empty")
	}

	methodsAvailable := []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut}
	if !slices.Contains(methodsAvailable, abd.Method) {
		return fmt.Errorf("invalid auth bearer dynamic method | Check auth.auth_bearer_dynamic.method | available options: %v", methodsAvailable)
	}

	if abd.Extract == "" {
		return errors.New("invalid auth bearer dynamic extract | Check auth.auth_bearer_dynamic.extract | value cannot be empty")
	}

	err := abd.EndpointConfig.Validate()
	if err != nil {
		return fmt.Errorf("invalid auth bearer dynamic endpoint config | Check auth.auth_bearer_dynamic.endpoint_config | %v", err.Error())
	}

	return nil
}

func (ab *S_AuthCookie) validate() error {
	if ab.Endpoint == "" {
		return errors.New("invalid auth cookie endpoint | Check auth.auth_cookie.endpoint | value cannot be empty")
	}

	methodsAvailable := []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut}
	if !slices.Contains(methodsAvailable, ab.Method) {
		return fmt.Errorf("invalid auth cookie method | Check auth.auth_cookie.method | available options: %v", methodsAvailable)
	}

	if len(ab.Extract) == 0 {
		return errors.New("invalid auth cookie extract | Check auth.auth_cookie.extract | value cannot be empty")
	}

	err := ab.EndpointConfig.Validate()
	if err != nil {
		return fmt.Errorf("invalid auth cookie endpoint config | Check auth.auth_cookie.endpoint_config | %v", err.Error())
	}

	return nil
}
