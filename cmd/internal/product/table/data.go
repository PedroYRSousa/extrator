package table

type tableColumnType string
type tableColumnFormat string
type tableColumnMode string

const (
	TABLE_COLUMN_TYPE_STRING   tableColumnType = "string"
	TABLE_COLUMN_TYPE_INTEGER  tableColumnType = "integer"
	TABLE_COLUMN_TYPE_DOUBLE   tableColumnType = "double"
	TABLE_COLUMN_TYPE_BOOLEAN  tableColumnType = "boolean"
	TABLE_COLUMN_TYPE_DATE     tableColumnType = "date"
	TABLE_COLUMN_TYPE_DATETIME tableColumnType = "datetime"
	TABLE_COLUMN_TYPE_ARRAY    tableColumnType = "array"

	TABLE_MODE_AUTO    tableColumnMode = "auto"
	TABLE_MODE_DEFINED tableColumnMode = "defined"

	TABLE_COLUMN_TYPE_DEFAULT            tableColumnType   = TABLE_COLUMN_TYPE_STRING
	TABLE_COLUMN_FORMAT_DATE_DEFAULT     tableColumnFormat = "2006-01-02"
	TABLE_COLUMN_FORMAT_DATETIME_DEFAULT tableColumnFormat = "2006-01-02T15:04:05Z07:00"
	TABLE_ROOT_DEFAULT                                     = "$"
	TABLE_COLUMNS_MODE_DEFAULT           tableColumnMode   = TABLE_MODE_AUTO
)

var (
	TYPES_AVAILABLE = []tableColumnType{TABLE_COLUMN_TYPE_STRING, TABLE_COLUMN_TYPE_INTEGER, TABLE_COLUMN_TYPE_DOUBLE, TABLE_COLUMN_TYPE_BOOLEAN, TABLE_COLUMN_TYPE_DATE, TABLE_COLUMN_TYPE_DATETIME, TABLE_COLUMN_TYPE_ARRAY}
	MODES_AVAILABLE = []tableColumnMode{TABLE_MODE_AUTO, TABLE_MODE_DEFINED}
)

type S_Column struct {
	Path  string `yaml:"path"`
	Alias string `yaml:"alias"`

	// Opcionais
	Type   *tableColumnType `yaml:"type,omitempty"`
	Format *tableColumnFormat
}

type S_Table struct {
	Name string `yaml:"name"`

	// Opcionais
	Root                  *string          `yaml:"root,omitempty"`
	ColumnsMode           *tableColumnMode `yaml:"columns_mode,omitempty"`
	Columns               *[]S_Column      `yaml:"columns,omitempty"`
	PrivateKeyColumnAlias *string          `yaml:"private_key_column_alias,omitempty"`
}
