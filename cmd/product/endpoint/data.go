package endpoint

type S_Endpoint struct {
	// Obrigatório
	Name        string
	Path        string
	Version     string `yaml:"version"`
	Description string `yaml:"description"`

	// Opcionais
}
