package page

import (
	"extrator/modules"
)

func (page *S_PaginatePage) format() *S_PaginatePage {
	modules.FormatString(page)

	return page
}
