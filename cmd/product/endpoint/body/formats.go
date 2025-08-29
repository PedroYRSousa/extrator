package body

import (
	"extrator/modules"
)

func (body *S_Body) format() *S_Body {
	modules.FormatString(body)

	return body
}
