package body

var (
	CONTENT_TYPE_ALLOWED = []string{"application/x-www-form-urlencoded", "multipart/form-data", "application/json", "application/xml", "application/graphql", "application/protobuf", "application/msgpack", "application/cbor", "application/x-ndjson", "text/plain", "text/csv", "application/octet-stream"}
)

type S_Body struct {
	ContentType string             `yaml:"content_type,omitempty" validate:"required,printascii"`
	File        *string            `yaml:"_file,omitempty" validate:"excluded_unless=Content nil FormFields nil MultiPart nil"`
	Content     *string            `yaml:"_content,omitempty" validate:"excluded_unless=File nil FormFields nil MultiPart nil"`
	FormFields  *map[string]string `yaml:"_form_fields,omitempty" validate:"excluded_unless=File nil Content nil MultiPart nil"`
	MultiPart   *struct {
		Fields map[string]string   `yaml:"fields" validate:"required"`
		Files  map[string][]string `yaml:"files" validate:"required"`
	} `yaml:"_multi_part,omitempty" validate:"excluded_unless=File nil FormFields nil Content nil"`
}
