package modules

import (
	"os"
	"path/filepath"
	"testing"
)

type sample struct {
	Name string `json:"name" yaml:"name" xml:"name"`
	Age  int    `json:"age" yaml:"age" xml:"age"`
}

// --- helpers ---

func writeTempFile(t *testing.T, content string, ext string) string {
	t.Helper()

	tmp := filepath.Join(t.TempDir(), "test."+ext)
	err := os.WriteFile(tmp, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	return tmp
}

// --- tests ---

func TestYamlToStruct_Success(t *testing.T) {
	content := `
name: John
age: 30
`
	path := writeTempFile(t, content, "yaml")

	var s sample
	err := YamlToStruct(path, &s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if s.Name != "John" || s.Age != 30 {
		t.Fatalf("unexpected struct values: %+v", s)
	}
}

func TestYamlToStruct_InvalidContent(t *testing.T) {
	path := writeTempFile(t, "invalid yaml", "yaml")

	var s sample
	err := YamlToStruct(path, &s)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestYamlToStruct_FileNotFound(t *testing.T) {
	var s sample
	err := YamlToStruct("nonexistent.yaml", &s)
	if err == nil {
		t.Fatal("expected error for nonexistent file, got nil")
	}
}

func TestJsonToStruct_Success(t *testing.T) {
	content := `{"name": "Alice", "age": 25}`
	path := writeTempFile(t, content, "json")

	var s sample
	err := JsonToStruct(path, &s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if s.Name != "Alice" || s.Age != 25 {
		t.Fatalf("unexpected struct values: %+v", s)
	}
}

func TestJsonToStruct_InvalidContent(t *testing.T) {
	path := writeTempFile(t, "{invalid json}", "json")

	var s sample
	err := JsonToStruct(path, &s)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestJsonToStruct_FileNotFound(t *testing.T) {
	var s sample
	err := JsonToStruct("nonexistent.json", &s)
	if err == nil {
		t.Fatal("expected error for nonexistent file, got nil")
	}
}

func TestXMLToStruct_Success(t *testing.T) {
	content := `<sample><name>Bob</name><age>40</age></sample>`
	path := writeTempFile(t, content, "xml")

	var s sample
	err := XMLToStruct(path, &s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if s.Name != "Bob" || s.Age != 40 {
		t.Fatalf("unexpected struct values: %+v", s)
	}
}

func TestXMLToStruct_InvalidContent(t *testing.T) {
	path := writeTempFile(t, "<invalid xml", "xml")

	var s sample
	err := XMLToStruct(path, &s)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestXMLToStruct_FileNotFound(t *testing.T) {
	var s sample
	err := XMLToStruct("nonexistent.xml", &s)
	if err == nil {
		t.Fatal("expected error for nonexistent file, got nil")
	}
}
