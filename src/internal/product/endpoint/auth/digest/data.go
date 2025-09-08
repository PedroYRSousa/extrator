package digest

type S_Digest struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}
