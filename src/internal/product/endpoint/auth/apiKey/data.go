package apikey

type S_ApiKey struct {
	Key      string `yaml:"key" validate:"required"`
	Value    string `yaml:"value" validate:"required"`
	Location string `yaml:"location" validate:"required,lowercase,oneof=header query cookie"`
}
