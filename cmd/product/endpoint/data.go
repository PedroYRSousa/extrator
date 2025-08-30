package endpoint

type S_Endpoint struct {
	Path string `validate:"required,filepath,printascii"`
	Name string `validate:"required,printascii"`

	Version     string `yaml:"version" validate:"required,printascii"`
	Description string `yaml:"description" validate:"required,printascii"`
}
