package client

import "net/http"

type S_Client struct {
	HttpClient *http.Client

	ProxyUrl                       *string `yaml:"_proxyUrl"`
	TimeoutInSeconds               int     `yaml:"_timeout_in_seconds" validate:"gte=0"`
	TimeoutDialInSeconds           int     `yaml:"_timeout_dial_in_seconds" validate:"gte=0"`
	KeepAliveDialInSeconds         int     `yaml:"_keep_alive_dial_in_seconds" validate:"gte=0"`
	SkipCertificate                bool    `yaml:"_skip_certificate"`
	TimeoutTLSHandshakeInSeconds   int     `yaml:"_timeout_tls_handshake_in_seconds" validate:"gte=0"`
	TimeoutExpectContinueInSeconds int     `yaml:"_timeout_expect_continue_in_seconds" validate:"gte=0"`
	MaxIdleConns                   int     `yaml:"_max_idle_conns" validate:"gte=0"`
	MaxIdleConnsPerHost            int     `yaml:"_max_idle_conns_per_host" validate:"gte=0"`
	TimeoutIdleConnInSeconds       int     `yaml:"_timeout_idle_conn_in_seconds" validate:"gte=0"`
	TimeoutResponseHeaderInSeconds int     `yaml:"_timeout_response_header_in_seconds" validate:"gte=0"`
}
