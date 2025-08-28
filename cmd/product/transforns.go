package product

import "extrator/modules"

func (product *S_Product) transform() error {
	err := modules.TransformString(product)
	if err != nil {
		return err
	}

	return nil
}
