package client

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"
)

func (c *S_Client) Mount() error {
	if c == nil {
		panic("Internal error | client.Mount")
	}

	c.HttpClient = &http.Client{}
	c.HttpClient.Timeout = time.Duration(c.TimeoutInSeconds) * time.Second // timeout total da requisição (DNS + conexão + resposta)

	transport := &http.Transport{}

	if c.ProxyUrl != nil {
		uri, err := url.Parse(*c.ProxyUrl)
		if err != nil {
			return err
		}
		transport.Proxy = http.ProxyURL(uri)
	}

	// Customização de conexão TCP
	transport.DialContext = (&net.Dialer{
		Timeout:   time.Duration(c.TimeoutInSeconds) * time.Second,       // timeout para abrir conexão
		KeepAlive: time.Duration(c.KeepAliveDialInSeconds) * time.Second, // tempo de keep-alive para conexões persistentes
	}).DialContext

	// Configuração TLS
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: c.SkipCertificate, // true = aceita certificado inválido (não recomendado)
	}
	transport.TLSHandshakeTimeout = time.Duration(c.TimeoutTLSHandshakeInSeconds) * time.Second
	transport.ExpectContinueTimeout = time.Duration(c.TimeoutExpectContinueInSeconds) * time.Second

	// Pool de conexões
	transport.MaxIdleConns = c.MaxIdleConns
	transport.MaxIdleConnsPerHost = c.MaxIdleConnsPerHost
	transport.IdleConnTimeout = time.Duration(c.TimeoutIdleConnInSeconds) * time.Second

	// Limites
	transport.ResponseHeaderTimeout = time.Duration(c.TimeoutResponseHeaderInSeconds) * time.Second

	c.HttpClient.Transport = transport

	return nil
}
