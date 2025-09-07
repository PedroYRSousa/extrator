package multipart

type S_Multipart struct {
	Fields *map[string]string
	Files  *[]map[string]map[string]string
}
