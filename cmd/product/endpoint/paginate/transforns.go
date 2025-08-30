package paginate

import "extrator/modules"

func (paginate *S_Paginate) transform() error {
	err := modules.TransformString(paginate)
	if err != nil {
		return err
	}

	return nil
}
