package page

func (page *S_PaginatePage) Check() error {
	if page == nil {
		return nil
	}

	page = page.format()

	err := page.validate()
	if err != nil {
		return err
	}

	err = page.transform()
	if err != nil {
		return err
	}

	return nil
}
