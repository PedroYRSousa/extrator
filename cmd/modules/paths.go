package modules

import (
	"os"
	"strings"
)

func GetDataPaths() string {
	return os.TempDir()
}

func JoinPaths(segments ...string) string {
	if len(segments) == 0 {
		return "."
	}

	parts := make([]string, 0, len(segments))
	absolute := false
	lastHadTrailingSlash := false

	// Encontrar último segmento não-vazio para checar barra final
	lastNonEmptyIdx := -1
	for i := len(segments) - 1; i >= 0; i-- {
		if segments[i] != "" {
			lastNonEmptyIdx = i
			break
		}
	}
	if lastNonEmptyIdx >= 0 && strings.HasSuffix(segments[lastNonEmptyIdx], "/") {
		lastHadTrailingSlash = true
	}

	for i, seg := range segments {
		if seg == "" {
			continue
		}

		// Se o segmento for absoluto, reinicia estado acumulado
		if strings.HasPrefix(seg, "/") {
			absolute = true
			parts = parts[:0]
		}

		// Remove barras iniciais e fatiar pelos separadores
		seg = strings.TrimLeft(seg, "/")
		chunks := strings.Split(seg, "/")

		for _, c := range chunks {
			if c == "" || c == "." {
				continue
			}
			if c == ".." {
				if len(parts) > 0 {
					parts = parts[:len(parts)-1]
				} else if !absolute {
					// Em relativo, preserva .. no início
					parts = append(parts, "..")
				}
				continue
			}
			parts = append(parts, c)
		}

		// Preservação de barra final só faz sentido no último segmento não-vazio
		if i != lastNonEmptyIdx {
			lastHadTrailingSlash = false
		}
	}

	var res string
	if absolute {
		res = "/" + strings.Join(parts, "/")
	} else {
		res = strings.Join(parts, "/")
	}

	if res == "" {
		if absolute {
			res = "/"
		} else {
			res = "."
		}
	} else if lastHadTrailingSlash && !strings.HasSuffix(res, "/") {
		// Preserva barra final se o último segmento não-vazio a possuía
		res += "/"
	}

	return res
}
