package endpoint_config

import "strings"

func (po *S_PaginationOffset) format() {
	po.Direction = paginationDirection(strings.TrimSpace(strings.ToLower(string(po.Direction))))
	po.Location = paginationLocation(strings.TrimSpace(strings.ToLower(string(po.Location))))
	po.Offset = strings.TrimSpace(po.Offset)
	po.Limit = strings.TrimSpace(po.Limit)
}

func (pp *S_PaginationPage) format() {
	pp.Direction = paginationDirection(strings.TrimSpace(strings.ToLower(string(pp.Direction))))
	pp.Location = paginationLocation(strings.TrimSpace(strings.ToLower(string(pp.Location))))
	pp.Page = strings.TrimSpace(pp.Page)
	if pp.PageSize <= 0 {
		pp.PageSize = 1
	}
}

func (pp *S_PaginationProperty) format() {
	pp.Property = strings.TrimSpace(pp.Property)
}

func (plh *S_PaginationLinkHeader) format() {
	plh.Header = strings.TrimSpace(plh.Header)
}

func (p *S_Pagination) format() {
	if p == nil {
		p = new(S_Pagination)
		*p = *newPagination()
		return
	}

	p.Mode = paginationMode(strings.TrimSpace(strings.ToLower(string(p.Mode))))
}

func (r *S_Retry) format() {
	if r == nil {
		r = new(S_Retry)
		*r = *newRetry()
		return
	}

	if r.Attempts == nil {
		r.Attempts = new(uint)
		*r.Attempts = ENDPOINT_RETRY_ATTEMPTS_DEFAULT
	}

	if r.DelayInSeconds == nil {
		r.DelayInSeconds = new(uint)
		*r.DelayInSeconds = ENDPOINT_RETRY_DELAY_IN_SECONDS_DEFAULT
	}
}

func (ec *S_EndpointConfig) format() {
	if ec == nil {
		ec = new(S_EndpointConfig)
		*ec = *newEndpointConfig()
		return
	}

	if ec.Method == nil {
		ec.Method = new(endpointMethod)
		*ec.Method = ENDPOINT_ENDPOINT_CONFIG_METHOD_DEFAULT
	}

	if ec.StatusCodeWaiting == nil {
		ec.StatusCodeWaiting = new([]uint)
		*ec.StatusCodeWaiting = ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_WAITING_DEFAULT
	}

	if ec.StatusCodeSuccess == nil {
		ec.StatusCodeSuccess = new([]uint)
		*ec.StatusCodeSuccess = ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SUCCESS_DEFAULT
	}

	if ec.StatusCodeSkip == nil {
		ec.StatusCodeSkip = new([]uint)
		*ec.StatusCodeSkip = ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SKIPPING_DEFAULT
	}

	if ec.WaitingTimeInSeconds == nil {
		ec.WaitingTimeInSeconds = new(uint)
		*ec.WaitingTimeInSeconds = ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_IN_SECONDS_DEFAULT
	}

	if ec.WaitingTimeErrorInSeconds == nil {
		ec.WaitingTimeErrorInSeconds = new(uint)
		*ec.WaitingTimeErrorInSeconds = ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_ERROR_IN_SECONDS_DEFAULT
	}

	if ec.TimeoutInSeconds == nil {
		ec.TimeoutInSeconds = new(uint)
		*ec.TimeoutInSeconds = ENDPOINT_ENDPOINT_CONFIG_TIMEOUT_IN_SECONDS_DEFAULT
	}
}
