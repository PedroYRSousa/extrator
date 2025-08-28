package endpoint

import "extrator/modules"

func (endpoint *S_Endpoint) transform() error {
	err := modules.TransformString(endpoint)
	if err != nil {
		return err
	}

	return nil
}
