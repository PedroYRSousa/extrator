package body

import "extrator/product/endpoint/endpoint_config/body/multi_part"

const (
	BODY_FORMAT_MULTIPART    = "multipart/form-data"
	BODY_FORMAT_URLENCODED   = "application/x-www-form-urlencoded"
	BODY_FORMAT_JSON         = "application/json"
	BODY_FORMAT_XML          = "application/xml"
	BODY_FORMAT_GRAPHQL      = "application/graphql"
	BODY_FORMAT_PROTOBUF     = "application/protobuf"
	BODY_FORMAT_MSGPACK      = "application/msgpack"
	BODY_FORMAT_CBOR         = "application/cbor"
	BODY_FORMAT_X_NDJSON     = "application/x-ndjson"
	BODY_FORMAT_PLAIN        = "text/plain"
	BODY_FORMAT_CSV          = "text/csv"
	BODY_FORMAT_OCTET_STREAM = "application/octet-stream"
)

var (
	BODY_FORMAT_ALLOWED = []string{BODY_FORMAT_MULTIPART, BODY_FORMAT_URLENCODED, BODY_FORMAT_JSON, BODY_FORMAT_XML, BODY_FORMAT_GRAPHQL, BODY_FORMAT_PROTOBUF, BODY_FORMAT_MSGPACK, BODY_FORMAT_CBOR, BODY_FORMAT_X_NDJSON, BODY_FORMAT_PLAIN, BODY_FORMAT_CSV, BODY_FORMAT_OCTET_STREAM}
)

type S_Body struct {
	ContentType string `yaml:"content_type" validate:"required,printascii"`

	File      *string                 `yaml:"_file,omitempty" validate:"excluded_unless=Content nil FormField nil MultiPart nil ContentType application/x-www-form-urlencoded ContentType multipart/form-data,printascii"`
	Content   *string                 `yaml:"_content,omitempty" validate:"excluded_unless=File nil FormField nil MultiPart nil ContentType application/x-www-form-urlencoded ContentType multipart/form-data,printascii"`
	FormField *map[string]string      `yaml:"_form_fields,omitempty" validate:"excluded_unless=File nil Content nil MultiPart nil,required_if=ContentType application/x-www-form-urlencoded"`
	MultiPart *multi_part.S_MultiPart `yaml:"_multi_part,omitempty" validate:"excluded_unless=File nil FormField nil Content nil,required_if=ContentType multipart/form-data"`
}
