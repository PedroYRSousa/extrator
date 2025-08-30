package property

func (paginateProperty *S_PaginateProperty) Check() error {
	if paginateProperty == nil {
		return nil
	}

	paginateProperty = paginateProperty.format()

	err := paginateProperty.validate()
	if err != nil {
		return err
	}

	err = paginateProperty.transform()
	if err != nil {
		return err
	}

	return nil
}
