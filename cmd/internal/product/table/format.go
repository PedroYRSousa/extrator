package table

import (
	"strings"
)

func (c *S_Column) format() {
	if c == nil {
		panic("Columns is required")
	}

	c.Path = strings.TrimSpace(c.Path)
	c.Alias = strings.TrimSpace(c.Alias)

	if c.Type == nil {
		c.Type = new(tableColumnType)
		*c.Type = TABLE_COLUMN_TYPE_DEFAULT
	} else {
		*c.Type = tableColumnType(strings.ToLower(string(*c.Type)))
	}

	switch *c.Type {
	case TABLE_COLUMN_TYPE_DATE:
		c.Format = new(tableColumnFormat)
		*c.Format = TABLE_COLUMN_FORMAT_DATE_DEFAULT
	case TABLE_COLUMN_TYPE_DATETIME:
		c.Format = new(tableColumnFormat)
		*c.Format = TABLE_COLUMN_FORMAT_DATETIME_DEFAULT
	}
}

func (t *S_Table) format() {
	t.Name = strings.TrimSpace(t.Name)

	if t.Root == nil {
		t.Root = new(string)
		*t.Root = TABLE_ROOT_DEFAULT
	} else {
		*t.Root = strings.TrimSpace(*t.Root)
	}

	if t.ColumnsMode == nil {
		t.ColumnsMode = new(tableColumnMode)
		*t.ColumnsMode = TABLE_COLUMNS_MODE_DEFAULT
	} else {
		*t.ColumnsMode = tableColumnMode(strings.TrimSpace(string(*t.ColumnsMode)))
	}
}
