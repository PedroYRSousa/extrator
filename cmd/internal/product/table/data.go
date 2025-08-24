package table

type S_Column struct {
	Path   string  `yaml:"path"`
	Alias  string  `yaml:"alias"`
	Type   string  `yaml:"type"`
	Format *string `yaml:"format,omitempty"`
}

type S_Table struct {
	Name                  string     `yaml:"name"`
	Root                  string     `yaml:"root"`
	ColumnsMode           string     `yaml:"columns_mode"`
	Columns               []S_Column `yaml:"columns"`
	PrivateKeyColumnAlias *string    `yaml:"private_key_column_alias,omitempty"`
}
