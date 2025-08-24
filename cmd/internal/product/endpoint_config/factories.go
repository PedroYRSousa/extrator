package endpoint_config

func newRetry() *S_Retry {
	attempts := ENDPOINT_RETRY_ATTEMPTS_DEFAULT
	delayInSeconds := ENDPOINT_RETRY_DELAY_IN_SECONDS_DEFAULT

	return &S_Retry{
		Attempts:       &attempts,
		DelayInSeconds: &delayInSeconds,
		BackoffFactor:  nil,
	}
}

func newEndpointConfig() *S_EndpointConfig {
	statusCodeWaiting := ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SKIPPING_DEFAULT
	statusCodeSuccess := ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SKIPPING_DEFAULT
	statusCodeSkip := ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SKIPPING_DEFAULT
	waitingTimeInSeconds := ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_IN_SECONDS_DEFAULT
	waitingTimeErrorInSeconds := ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_ERROR_IN_SECONDS_DEFAULT
	timeoutInSeconds := ENDPOINT_ENDPOINT_CONFIG_TIMEOUT_IN_SECONDS_DEFAULT

	return &S_EndpointConfig{
		StatusCodeWaiting:         &statusCodeWaiting,
		StatusCodeSuccess:         &statusCodeSuccess,
		StatusCodeSkip:            &statusCodeSkip,
		WaitingTimeInSeconds:      &waitingTimeInSeconds,
		WaitingTimeErrorInSeconds: &waitingTimeErrorInSeconds,
		TimeoutInSeconds:          &timeoutInSeconds,
		Headers:                   nil,
		Body:                      nil,
		QueryParams:               nil,
		Retry:                     newRetry(),
	}
}
