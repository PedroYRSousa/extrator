package product

func (product *S_Product) check() error {
	product = product.format()

	err := product.validate()
	if err != nil {
		return err
	}

	err = product.transform()
	if err != nil {
		return err
	}

	return nil
}
