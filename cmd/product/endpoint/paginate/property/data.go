package property

type S_PaginateProperty struct {
	Property string `yaml:"property,omitempty" validate:"required,printascii"`
}
