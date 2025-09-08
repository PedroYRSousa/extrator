package multipart

type S_Multipart struct {
	Fields []map[string]string            `yaml:"fields"`
	Files  []map[string]map[string]string `yaml:"files"`
}
