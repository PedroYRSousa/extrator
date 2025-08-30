package link_header

import "extrator/modules"

func (paginateLinkHeader *S_PaginateLinkHeader) transform() error {
	err := modules.TransformString(paginateLinkHeader)
	if err != nil {
		return err
	}

	return nil
}
