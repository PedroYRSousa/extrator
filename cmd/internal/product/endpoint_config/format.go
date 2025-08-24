package endpointconfig

func (r *S_Retry) format() {
	if r == nil {
		r = new(S_Retry)
		*r = *newRetry()
		return
	}

	if r.Attempts == nil {
		r.Attempts = new(int)
		*r.Attempts = ENDPOINT_RETRY_ATTEMPTS_DEFAULT
	}

	if r.DelayInSeconds == nil {
		r.DelayInSeconds = new(int)
		*r.DelayInSeconds = ENDPOINT_RETRY_DELAY_IN_SECONDS_DEFAULT
	}
}

func (ec *S_EndpointConfig) format() {
	if ec == nil {
		ec = new(S_EndpointConfig)
		*ec = *newEndpointConfig()
		return
	}

	if ec.StatusCodeWaiting == nil {
		ec.StatusCodeWaiting = new([]int)
		*ec.StatusCodeWaiting = ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_WAITING_DEFAULT
	}

	if ec.StatusCodeSuccess == nil {
		ec.StatusCodeSuccess = new([]int)
		*ec.StatusCodeSuccess = ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SUCCESS_DEFAULT
	}

	if ec.StatusCodeSkip == nil {
		ec.StatusCodeSkip = new([]int)
		*ec.StatusCodeSkip = ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SKIPPING_DEFAULT
	}

	if ec.WaitingTimeInSeconds == nil {
		ec.WaitingTimeInSeconds = new(int)
		*ec.WaitingTimeInSeconds = ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_IN_SECONDS_DEFAULT
	}

	if ec.WaitingTimeErrorInSeconds == nil {
		ec.WaitingTimeErrorInSeconds = new(int)
		*ec.WaitingTimeErrorInSeconds = ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_ERROR_IN_SECONDS_DEFAULT
	}

	if ec.TimeoutInSeconds == nil {
		ec.TimeoutInSeconds = new(int)
		*ec.TimeoutInSeconds = ENDPOINT_ENDPOINT_CONFIG_TIMEOUT_IN_SECONDS_DEFAULT
	}

	ec.Retry.format()
}
