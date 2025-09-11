package client

import "net/http"

type S_Client struct {
	HttpClient *http.Client

	ProxyUrl                       *string `yaml:"_proxyUrl" validate:"required,url,printascii"`
	TimeoutInSeconds               int     `yaml:"_timeout_in_seconds"`
	TimeoutDialInSeconds           int     `yaml:"_timeout_dial_in_seconds"`
	KeepAliveDialInSeconds         int     `yaml:"_keep_alive_dial_in_seconds"`
	SkipCertificate                bool    `yaml:"_skip_certificate"`
	TimeoutTLSHandshakeInSeconds   int     `yaml:"_timeout_tls_handshake_in_seconds"`
	TimeoutExpectContinueInSeconds int     `yaml:"_timeout_expect_continue_in_seconds"`
	MaxIdleConns                   int     `yaml:"_max_idle_conns"`
	MaxIdleConnsPerHost            int     `yaml:"_max_idle_conns_per_host"`
	TimeoutIdleConnInSeconds       int     `yaml:"_timeout_idle_conn_in_seconds"`
	TimeoutResponseHeaderInSeconds int     `yaml:"_timeout_response_header_in_seconds"`
}
