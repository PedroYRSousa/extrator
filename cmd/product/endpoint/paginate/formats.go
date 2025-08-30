package paginate

import (
	"extrator/modules"
)

func (paginate *S_Paginate) format() *S_Paginate {
	modules.FormatString(paginate)

	return paginate
}
