package table

import (
	"errors"
	"fmt"
	"slices"
)

func (c *S_Column) validate() error {
	c.format()

	if c.Path == "" {
		return errors.New("check tables.columns.path | path cannot be empty")
	}

	if c.Alias == "" {
		return errors.New("check tables.columns.alias | alias cannot be empty")
	}

	if !slices.Contains(TYPES_AVAILABLE, *c.Type) {
		return fmt.Errorf("check tables.columns.type | available options: %v", TYPES_AVAILABLE)
	}

	return nil
}

func (t *S_Table) Validate() error {
	t.format()

	if t.Name == "" {
		return errors.New("check tables.name | name cannot be empty")
	}

	if !slices.Contains(MODES_AVAILABLE, *t.ColumnsMode) {
		return fmt.Errorf("check tables.columns_mode | available options: %v", MODES_AVAILABLE)
	}

	if *t.ColumnsMode != TABLE_MODE_AUTO {
		if *t.Root == "" || (*t.Root)[0] != '$' {
			return errors.New("check tables.root | root cannot be empty and must start with '$' for JSONPath format")
		}

		if t.PrivateKeyColumnAlias == nil || *t.PrivateKeyColumnAlias == "" {
			return errors.New("check tables.private_key_column_alias | private_key_column_alias cannot be empty if columns_mode is not auto")
		}

		if t.Columns == nil || len(*t.Columns) <= 0 {
			return errors.New("check tables.columns | columns cannot be empty if columns_mode is not auto")
		}

		foundPrivateKeyColumn := false
		for _, column := range *t.Columns {
			if t.PrivateKeyColumnAlias != nil && column.Alias == *t.PrivateKeyColumnAlias {
				foundPrivateKeyColumn = true
			}

			err := column.validate()
			if err != nil {
				return err
			}
		}

		if t.PrivateKeyColumnAlias != nil && !foundPrivateKeyColumn {
			return errors.New("check tables.private_key_column_alias | alias not found in columns")
		}
	}

	return nil
}
