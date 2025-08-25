package table

func (c *S_Column) transform() error {
	return nil
}

func (t *S_Table) Transform() error {
	if *t.ColumnsMode != TABLE_MODE_AUTO {
		for _, column := range *t.Columns {
			err := column.transform()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
