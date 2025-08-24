package table

import (
	"errors"
	"fmt"
	"slices"
)

func (t *S_Table) Validate() error {
	if t.Name == "" {
		return errors.New("invalid table name | Check tables.name | name cannot be empty")
	}

	if (t.Root == "" || t.Root == "/") && t.Root[0] != '$' {
		return errors.New("invalid table root | Check tables.root | root cannot be empty or '/' when response format is json or xml")
	}

	columnsModesAvailable := []string{"auto", "defined"}
	if !slices.Contains(columnsModesAvailable, t.ColumnsMode) {
		return fmt.Errorf("invalid table columns mode | Check tables.columns_mode | available options: %v", columnsModesAvailable)
	}

	foundPrivateKeyColumn := false

	for _, column := range t.Columns {
		if t.PrivateKeyColumnAlias != nil && column.Alias == *t.PrivateKeyColumnAlias {
			foundPrivateKeyColumn = true
		}

		err := column.validate()
		if err != nil {
			return err
		}
	}

	if t.PrivateKeyColumnAlias != nil && !foundPrivateKeyColumn {
		return errors.New("invalid table private key column alias | Check tables.private_key_column_alias | alias not found in columns")
	}

	return nil
}

func (c *S_Column) validate() error {
	if c.Path == "" {
		return errors.New("invalid column path | Check tables.columns.path | path cannot be empty")
	}

	if c.Alias == "" {
		return errors.New("invalid column alias | Check tables.columns.alias | alias cannot be empty")
	}

	availableTypes := []string{"string", "integer", "double", "bool", "date", "datetime"}
	if !slices.Contains(availableTypes, c.Type) {
		return fmt.Errorf("invalid column type | Check tables.columns.type | available options: %v", availableTypes)
	}

	if (c.Type == "date" || c.Type == "datetime") && (c.Format == nil || *c.Format == "") {
		return errors.New("invalid column format | Check tables.columns.format | format cannot be empty when type is 'date' or 'datetime'")
	}

	return nil
}
