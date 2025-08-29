package body

var (
	CONTENT_TYPE_ALLOWED = []string{"application/json", "application/xml", "application/graphql", "application/protobuf", "application/msgpack", "application/cbor", "application/x-ndjson", "text/plain", "text/csv", "application/octet-stream"}
)

type S_Body struct {
	ContentType string             `yaml:"content_type,omitempty" validate:"required,printascii"`
	File        *string            `yaml:"_file,omitempty" validate:"excluded_unless=Content nil FormFields nil MultiPart nil"`
	Content     *string            `yaml:"_content,omitempty" validate:"excluded_unless=File nil FormFields nil MultiPart nil"`
	FormFields  *map[string]string `yaml:"_form_fields,omitempty" validate:"omitempty,required_if=ContentType application/x-www-form-urlencoded"`
	MultiPart   *struct {
		fields map[string]string   `yaml:"fields,omitempty" validate:"required,printascii"`
		files  map[string][]string `yaml:"files,omitempty" validate:"required,printascii"`
	} `yaml:"_multi_part,omitempty" validate:"omitempty,required_if=ContentType multipart/form-data"`
}
