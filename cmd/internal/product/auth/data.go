package auth

import (
	endpointconfig "extrator/internal/product/endpoint_config"
	"net/http"
)

type authMode string
type authApiKeyLocation string

const (
	AUTH_MODE_NONE           authMode = "none"
	AUTH_MODE_BASIC          authMode = "basic"
	AUTH_MODE_BEARER         authMode = "bearer"
	AUTH_MODE_BEARER_DYNAMIC authMode = "bearer_dynamic"
	AUTH_MODE_COOKIE         authMode = "cookie"
	AUTH_MODE_API_KEY        authMode = "api_key"

	AUTH_API_KEY_LOCATION_HEADER      authApiKeyLocation = "header"
	AUTH_API_KEY_LOCATION_QUERY_PARAM authApiKeyLocation = "query_param"
)

var (
	MODES_AVAILABLE     = []authMode{AUTH_MODE_NONE, AUTH_MODE_BASIC, AUTH_MODE_BEARER, AUTH_MODE_BEARER_DYNAMIC, AUTH_MODE_COOKIE, AUTH_MODE_API_KEY}
	AVAILABLE_LOCATIONS = []authApiKeyLocation{AUTH_API_KEY_LOCATION_HEADER, AUTH_API_KEY_LOCATION_QUERY_PARAM}
	METHODS_AVAILABLE   = []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut}
)

type S_Basic struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type S_Bearer struct {
	Name   string `yaml:"name"`
	Prefix string `yaml:"prefix"`
	Token  string `yaml:"token"`
}

type S_ApiKey struct {
	Name     string             `yaml:"name"`
	Location authApiKeyLocation `yaml:"location"`
	Value    string             `yaml:"value"`
}

type S_Cookie struct {
	Endpoint string   `yaml:"endpoint"`
	Method   string   `yaml:"method"`
	Extract  []string `yaml:"extract"`

	// Opcionais
	EndpointConfig *endpointconfig.S_EndpointConfig `yaml:"endpoint_config"`
}

type S_BearerDynamic struct {
	Name     string `yaml:"name"`
	Prefix   string `yaml:"prefix"`
	Endpoint string `yaml:"endpoint"`
	Method   string `yaml:"method"`
	Extract  string `yaml:"extract"`

	// Opcionais
	EndpointConfig *endpointconfig.S_EndpointConfig `yaml:"endpoint_config"`
}

type S_Auth struct {
	Mode authMode `yaml:"mode"`

	// Opcionais
	Basic         *S_Basic         `yaml:"basic,omitempty"`
	Bearer        *S_Bearer        `yaml:"bearer,omitempty"`
	ApiKey        *S_ApiKey        `yaml:"api_key,omitempty"`
	BearerDynamic *S_BearerDynamic `yaml:"bearer_dynamic,omitempty"`
	Cookie        *S_Cookie        `yaml:"cookie,omitempty"`
}
