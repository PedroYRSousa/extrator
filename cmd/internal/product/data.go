package product

import (
	"extrator/internal/product/auth"
	"extrator/internal/product/endpoint"
	"extrator/internal/product/pagination"
	"extrator/internal/product/table"
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

type S_ProductEndpoint struct {
	ProductName string
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`

	Endpoint   endpoint.S_Endpoint     `yaml:"endpoint"`
	Pagination pagination.S_Pagination `yaml:"pagination"`
	Auth       auth.S_Auth             `yaml:"auth"`
	Tables     []table.S_Table         `yaml:"tables"`
}
