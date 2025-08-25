package http_request

import "errors"

var (
	ErrStatusCodeWaiting = errors.New("waiting status code reached")
	ErrStatusCodeSkip    = errors.New("skip status code reached")
	ErrStatusCodeError   = errors.New("error status code reached")
)
