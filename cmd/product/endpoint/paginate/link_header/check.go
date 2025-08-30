package link_header

func (paginateLinkHeader *S_PaginateLinkHeader) Check() error {
	if paginateLinkHeader == nil {
		return nil
	}

	paginateLinkHeader = paginateLinkHeader.format()

	err := paginateLinkHeader.validate()
	if err != nil {
		return err
	}

	err = paginateLinkHeader.transform()
	if err != nil {
		return err
	}

	return nil
}
