package paginate

func (paginate *S_Paginate) Check() error {
	if paginate == nil {
		return nil
	}

	paginate = paginate.format()

	err := paginate.Page.Check()
	if err != nil {
		return err
	}

	err = paginate.Property.Check()
	if err != nil {
		return err
	}

	err = paginate.LinkHeader.Check()
	if err != nil {
		return err
	}

	err = paginate.validate()
	if err != nil {
		return err
	}

	err = paginate.transform()
	if err != nil {
		return err
	}

	return nil
}
