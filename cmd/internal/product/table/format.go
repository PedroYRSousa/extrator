package table

import (
	"strings"
)

func (t *S_Table) Format() {
	t.ColumnsMode = strings.ToLower(t.ColumnsMode)
	for index := range t.Columns {
		t.Columns[index].format()
	}
}

func (c *S_Column) format() {
	c.Type = strings.ToLower(c.Type)
}
