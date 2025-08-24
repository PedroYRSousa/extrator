package product

import (
	"extrator/internal/product/auth"
	"extrator/internal/product/endpoint"
	"extrator/internal/product/tables"
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

type S_Product struct {
	Name        string `yaml:"name"`
	Path        string
	Version     string `yaml:"version"`
	Description string `yaml:"description"`

	Endpoint endpoint.S_Endpoint `yaml:"endpoint"`

	Auth auth.S_Auth `yaml:"auth"`

	Tables []tables.S_Table `yaml:"tables"`
}
