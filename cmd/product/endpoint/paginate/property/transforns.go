package property

import "extrator/modules"

func (paginateProperty *S_PaginateProperty) transform() error {
	err := modules.TransformString(paginateProperty)
	if err != nil {
		return err
	}

	return nil
}
