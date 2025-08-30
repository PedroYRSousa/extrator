package property

import (
	"extrator/modules"
)

func (paginateProperty *S_PaginateProperty) format() *S_PaginateProperty {
	modules.FormatString(paginateProperty)

	return paginateProperty
}
