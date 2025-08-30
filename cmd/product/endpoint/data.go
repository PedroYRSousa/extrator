package endpoint

import (
	"extrator/product/endpoint/body"
	"extrator/product/endpoint/paginate"
	"net/http"
)

const (
	RESPONSE_FORMAT_JSON        = "json"
	RESPONSE_FORMAT_XML         = "xml"
	RESPONSE_FORMAT_ZIPPED_JSON = "zip/json"
	RESPONSE_FORMAT_ZIPPED_XML  = "zip/xml"

	METHOD_GET  = http.MethodGet
	METHOD_POST = http.MethodPost

	DEFAULT_METHOD          = METHOD_GET
	DEFAULT_RESPONSE_FORMAT = RESPONSE_FORMAT_JSON
)

var (
	METHODS_ALLOWED         = []string{METHOD_GET, METHOD_POST}
	RESPONSE_FORMAT_ALLOWED = []string{RESPONSE_FORMAT_JSON, RESPONSE_FORMAT_XML, RESPONSE_FORMAT_ZIPPED_JSON, RESPONSE_FORMAT_ZIPPED_XML}
)

type S_Endpoint struct {
	Request        *http.Request
	Name           string               `validate:"required,printascii"`
	Path           string               `validate:"required,file,printascii"`
	Version        string               `yaml:"version" validate:"required,printascii"`
	Description    string               `yaml:"description" validate:"required,printascii,min=15"`
	URL            string               `yaml:"url" validate:"required,printascii,url"`
	Method         *string              `yaml:"_method,omitempty" validate:"printascii"`
	ResponseFormat *string              `yaml:"_response_format,omitempty" validate:"printascii"`
	QueryParams    *map[string]string   `yaml:"_query_params,omitempty" validate:"printascii"`
	Headers        *map[string]string   `yaml:"_headers,omitempty" validate:"printascii"`
	Body           *body.S_Body         `yaml:"_body,omitempty"`
	Paginate       *paginate.S_Paginate `yaml:"_paginate,omitempty"`
}
