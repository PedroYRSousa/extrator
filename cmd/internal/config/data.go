package config

import (
	"extrator/internal/config/auth"
	"extrator/internal/config/endpoint"
	"extrator/internal/config/tables"
)

const configPath = "../config"

type s_EndpointConfigFile struct {
	Name string
	Path string
}

type s_ProductConfigFile struct {
	Name      string
	Path      string
	Endpoints []s_EndpointConfigFile
}

type s_ConfigFile struct {
	products []s_ProductConfigFile
}

type S_Config struct {
	Name        string `yaml:"name" json:"name"`
	Version     string `yaml:"version" json:"version"`
	Description string `yaml:"description" json:"description"`

	Endpoint endpoint.S_Endpoint `yaml:"endpoint" json:"endpoint"`

	Auth auth.S_Auth `yaml:"auth" json:"auth"`

	Tables []tables.S_Table `yaml:"tables" json:"tables"`
}
