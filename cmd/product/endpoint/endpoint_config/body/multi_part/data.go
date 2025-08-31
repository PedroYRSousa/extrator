package multi_part

type S_MultiPart struct {
	Fields map[string]string   `yaml:"fields" validate:"required,printascii"`
	Files  map[string][]string `yaml:"files" validate:"required,printascii"`
}
