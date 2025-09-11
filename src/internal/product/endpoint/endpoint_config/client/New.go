package client

import "net/http"

func New() *S_Client {
	return &S_Client{
		HttpClient:                     http.DefaultClient,
		ProxyUrl:                       nil,
		TimeoutInSeconds:               0,
		TimeoutDialInSeconds:           0,
		KeepAliveDialInSeconds:         0,
		SkipCertificate:                false,
		TimeoutTLSHandshakeInSeconds:   0,
		TimeoutExpectContinueInSeconds: 0,
		MaxIdleConns:                   0,
		MaxIdleConnsPerHost:            0,
		TimeoutIdleConnInSeconds:       0,
		TimeoutResponseHeaderInSeconds: 0,
	}
}
