package auth

import "extrator/internal/product/endpoint"

type S_AuthBasic struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type S_AuthBearer struct {
	Name   string `yaml:"name"`
	Prefix string `yaml:"prefix"`
	Token  string `yaml:"token"`
}

type S_AuthApiKey struct {
	Name     string `yaml:"name"`
	Location string `yaml:"location"`
	Value    string `yaml:"value"`
}

type S_AuthCookie struct {
	Endpoint       string                    `yaml:"endpoint"`
	Method         string                    `yaml:"method"`
	Extract        []string                  `yaml:"extract"`
	EndpointConfig endpoint.S_EndpointConfig `yaml:"endpoint_config"`
}

type S_AuthBearerDynamic struct {
	Name           string                    `yaml:"name"`
	Prefix         string                    `yaml:"prefix"`
	Endpoint       string                    `yaml:"endpoint"`
	Method         string                    `yaml:"method"`
	Extract        string                    `yaml:"extract"`
	EndpointConfig endpoint.S_EndpointConfig `yaml:"endpoint_config"`
}

type S_Auth struct {
	AuthMode          string               `yaml:"auth_mode"`
	AuthBasic         *S_AuthBasic         `yaml:"auth_basic,omitempty"`
	AuthBearer        *S_AuthBearer        `yaml:"auth_bearer,omitempty"`
	AuthApiKey        *S_AuthApiKey        `yaml:"auth_api_key,omitempty"`
	AuthBearerDynamic *S_AuthBearerDynamic `yaml:"auth_bearer_dynamic,omitempty"`
	AuthCookie        *S_AuthCookie        `yaml:"auth_cookie,omitempty"`
}
