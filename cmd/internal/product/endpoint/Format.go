package endpoint

import (
	"strings"
)

func (e *S_Endpoint) Format() {
	e.Method = strings.TrimSpace(strings.ToUpper(e.Method))
	e.URI = strings.TrimSpace(strings.ToLower(e.URI))
	e.ResponseFormat = strings.TrimSpace(strings.ToLower(e.ResponseFormat))
	if e.LimitExtract == nil {
		e.LimitExtract = new(int)
		*e.LimitExtract = -1
	}
	e.EndpointConfig.format()
}

func (qp *S_QueryParam) format() {
	qp.Name = strings.TrimSpace(qp.Name)
}

func (r *S_Retry) format() {
	if r.BackoffFactor == nil {
		r.BackoffFactor = new(int)
		*r.BackoffFactor = 0
	}
}

func (p *S_Pagination) format() {
	if p.Mode == nil {
		p.Mode = new(string)
		*p.Mode = "none"
	}
	*p.Mode = strings.TrimSpace(strings.ToLower(*p.Mode))
	p.Offset = strings.TrimSpace(p.Offset)
	p.Limit = strings.TrimSpace(p.Limit)
	p.Page = strings.TrimSpace(p.Page)
	p.Property = strings.TrimSpace(p.Property)
	p.Header = strings.TrimSpace(p.Header)
	p.Direction = strings.ToLower(p.Direction)
	p.Location = strings.ToLower(p.Location)
}

func (ec *S_EndpointConfig) format() {
	for index := range *ec.QueryParams {
		(*ec.QueryParams)[index].format()
	}
	ec.Retry.format()
	ec.Pagination.format()
}
