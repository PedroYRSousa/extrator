package body

import "extrator/modules"

func (body *S_Body) transform() error {
	err := modules.TransformString(body)
	if err != nil {
		return err
	}

	return nil
}
