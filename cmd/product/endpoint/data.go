package endpoint

type S_Endpoint struct {
	Name        string `validate:"required,printascii"`
	Path        string `validate:"required,file,printascii"`
	Version     string `yaml:"version" validate:"required,printascii"`
	Description string `yaml:"description" validate:"required,printascii,min=15"`
}
