package modules

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// --- tests ---

func TestMakeRequest_Success(t *testing.T) {
	// Cria servidor de teste que retorna status 200 e corpo "ok"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	res, body, err := MakeRequest(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	if string(body) != "ok" {
		t.Fatalf("expected body 'ok', got %q", string(body))
	}
}

func TestMakeRequest_ServerError(t *testing.T) {
	// Servidor retorna 500 com corpo
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)

	res, body, err := MakeRequest(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", res.StatusCode)
	}

	if !strings.Contains(string(body), "internal error") {
		t.Fatalf("expected body to contain 'internal error', got %q", string(body))
	}
}

func TestMakeRequest_InvalidURL(t *testing.T) {
	// Cria request com URL inválida (simula erro no client.Do)
	req, _ := http.NewRequest(http.MethodGet, "http://invalid.invalid", nil)

	// Força para não esperar tempo de DNS
	client := &http.Client{
		Transport: roundTripperFunc(func(*http.Request) (*http.Response, error) {
			return nil, errors.New("network error")
		}),
	}

	http.DefaultClient = client // sobrescreve temporariamente o DefaultClient

	_, _, err := MakeRequest(req)
	if err == nil || !strings.Contains(err.Error(), "network error") {
		t.Fatalf("expected network error, got %v", err)
	}
}

// --- helper para mockar client.Do ---
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
