package modules

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

// --- mapTranslate tests ---

func TestMapTranslate_KnownTags(t *testing.T) {
	tests := []struct {
		tag      string
		expected string
	}{
		{"dirpath", "The field must be a valid directory path"},
		{"filepath", "The field must be a valid file path"},
	}

	for _, tt := range tests {
		got := mapTranslate(tt.tag)
		if got != tt.expected {
			t.Fatalf("tag %q: expected %q, got %q", tt.tag, tt.expected, got)
		}
	}
}

func TestMapTranslate_UnknownTag(t *testing.T) {
	tag := "required"
	got := mapTranslate(tag)
	if got != tag {
		t.Fatalf("expected %q, got %q", tag, got)
	}
}

// --- TranslateValidatorError tests ---

type sampleDirPath struct {
	Dir string `validate:"dirpath"`
	Req string `validate:"required"`
}

type sampleFilePath struct {
	File string `validate:"file"`
	Req  string `validate:"required"`
}

func TestTranslateValidatorError_DirPath(t *testing.T) {
	validate := validator.New()

	s := sampleDirPath{Dir: "sample.Dir", Req: "ok"}
	err := validate.Struct(s)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	// Captura os erros do validador
	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		t.Fatalf("expected ValidationErrors, got %T", err)
	}

	trErr := TranslateValidatorError(verrs)
	if trErr == nil {
		t.Fatal("expected translated error, got nil")
	}

	msg := trErr.Error()
	if !strings.Contains(msg, "validation Error on field (sampleFilePath.Dir)") {
		t.Fatalf("unexpected message: %v", msg)
	}
	if !strings.Contains(msg, "The field must be a valid directory path") {
		t.Fatalf("expected dirpath message, got: %v", msg)
	}
}

func TestTranslateValidatorError_FilePath(t *testing.T) {
	validate := validator.New()

	s := sampleFilePath{File: "some.file", Req: "ok"}
	err := validate.Struct(s)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	// Captura os erros do validador
	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		t.Fatalf("expected ValidationErrors, got %T", err)
	}

	trErr := TranslateValidatorError(verrs)
	if trErr == nil {
		t.Fatal("expected translated error, got nil")
	}

	msg := trErr.Error()
	if !strings.Contains(msg, "validation Error on field (sampleFilePath.File)") {
		t.Fatalf("unexpected message: %v", msg)
	}
	if !strings.Contains(msg, "The field must be a valid file path") {
		t.Fatalf("expected dirpath message, got: %v", msg)
	}
}

func TestTranslateValidatorError_Required(t *testing.T) {
	validate := validator.New()

	s := sampleDirPath{Dir: ".", Req: ""}
	err := validate.Struct(s)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	verrs := err.(validator.ValidationErrors)
	trErr := TranslateValidatorError(verrs)
	if trErr == nil {
		t.Fatal("expected translated error, got nil")
	}

	// Para required, deve retornar o nome da tag
	msg := trErr.Error()
	if !strings.Contains(msg, "required") {
		t.Fatalf("expected 'required' in message, got: %v", msg)
	}
}

func TestTranslateValidatorError_NoErrors(t *testing.T) {
	var verrs validator.ValidationErrors // nil slice
	trErr := TranslateValidatorError(verrs)

	if trErr != nil {
		t.Fatalf("expected nil, got %v", trErr)
	}
}
