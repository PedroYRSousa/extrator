package auth

import (
	endpointconfig "extrator/internal/product/endpoint_config"
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
	Name     string `yaml:"name"`
	Location string `yaml:"location"`
	Value    string `yaml:"value"`
}

type S_Cookie struct {
	Endpoint       string                          `yaml:"endpoint"`
	Method         string                          `yaml:"method"`
	Extract        []string                        `yaml:"extract"`
	EndpointConfig endpointconfig.S_EndpointConfig `yaml:"endpoint_config"`
}

type S_BearerDynamic struct {
	Name           string                          `yaml:"name"`
	Prefix         string                          `yaml:"prefix"`
	Endpoint       string                          `yaml:"endpoint"`
	Method         string                          `yaml:"method"`
	Extract        string                          `yaml:"extract"`
	EndpointConfig endpointconfig.S_EndpointConfig `yaml:"endpoint_config"`
}

type S_Auth struct {
	Mode          string           `yaml:"mode"`
	Basic         *S_Basic         `yaml:"basic,omitempty"`
	Bearer        *S_Bearer        `yaml:"bearer,omitempty"`
	ApiKey        *S_ApiKey        `yaml:"api_key,omitempty"`
	BearerDynamic *S_BearerDynamic `yaml:"bearer_dynamic,omitempty"`
	Cookie        *S_Cookie        `yaml:"cookie,omitempty"`
}
