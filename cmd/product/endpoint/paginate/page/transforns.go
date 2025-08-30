package page

import "extrator/modules"

func (page *S_PaginatePage) transform() error {
	err := modules.TransformString(page)
	if err != nil {
		return err
	}

	return nil
}
