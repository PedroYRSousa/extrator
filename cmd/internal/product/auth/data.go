package auth

import (
	endpointconfig "extrator/internal/product/endpoint_config"
	"net/http"
)

type authMode string
type authApiKeyLocation string

const (
	AUTH_MODE_NONE    authMode = "none"
	AUTH_MODE_BASIC   authMode = "basic"
	AUTH_MODE_BEARER  authMode = "bearer"
	AUTH_MODE_COOKIE  authMode = "cookie"
	AUTH_MODE_API_KEY authMode = "api_key"

	AUTH_API_KEY_LOCATION_HEADER      authApiKeyLocation = "header"
	AUTH_API_KEY_LOCATION_QUERY_PARAM authApiKeyLocation = "query_param"

	AUTH_BEARER_NAME_DEFAULT      = "Authorization"
	AUTH_BEARER_PREFIX_DEFAULT    = "Bearer"
	AUTH_BEARER_DYNAMIC_DEFAULT   = false
	AUTH_API_KEY_LOCATION_DEFAULT = AUTH_API_KEY_LOCATION_HEADER
)

var (
	MODES_AVAILABLE     = []authMode{AUTH_MODE_NONE, AUTH_MODE_BASIC, AUTH_MODE_BEARER, AUTH_MODE_COOKIE, AUTH_MODE_API_KEY}
	AVAILABLE_LOCATIONS = []authApiKeyLocation{AUTH_API_KEY_LOCATION_HEADER, AUTH_API_KEY_LOCATION_QUERY_PARAM}
	METHODS_AVAILABLE   = []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut}
)

type S_Basic struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type S_DynamicDetails struct {
	Endpoint string `yaml:"endpoint"`
	Extract  string `yaml:"extract"`

	// Opcionais
	EndpointConfig *endpointconfig.S_EndpointConfig `yaml:"endpoint_config"`
}

type S_Bearer struct {
	Token string `yaml:"token"`

	// Opcionais
	Dynamic        *bool             `yaml:"dynamic,omitempty"`
	Name           *string           `yaml:"name,omitempty"`
	Prefix         *string           `yaml:"prefix,omitempty"`
	DynamicDetails *S_DynamicDetails `yaml:"dynamic_details,omitempty"`
}

type S_ApiKey struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`

	// Opcionais
	Location *authApiKeyLocation `yaml:"location,omitempty"`
}

type S_Cookie struct {
	Endpoint string `yaml:"endpoint"`

	// Opcionais
	EndpointConfig *endpointconfig.S_EndpointConfig `yaml:"endpoint_config"`
}

type S_Auth struct {
	Mode authMode `yaml:"mode"`

	// Opcionais
	Basic  *S_Basic  `yaml:"basic,omitempty"`
	Bearer *S_Bearer `yaml:"bearer,omitempty"`
	ApiKey *S_ApiKey `yaml:"api_key,omitempty"`
	Cookie *S_Cookie `yaml:"cookie,omitempty"`
}
