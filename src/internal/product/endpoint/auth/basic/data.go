package basic

type S_Basic struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}
