package modules

import (
	"errors"
	"strings"
	"testing"
	"time"
)

// --- helpers ---

// mockFunc retorna uma função que falha N vezes antes de retornar sucesso.
func mockFunc(failTimes int, err error) func() error {
	attempt := 0
	return func() error {
		if attempt < failTimes {
			attempt++
			return err
		}
		return nil
	}
}

// --- tests ---

func TestExecWithAttempts_SuccessFirstTry(t *testing.T) {
	err := ExecWithAttempts(3, 1, func() error {
		return nil
	})

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestExecWithAttempts_SuccessAfterRetries(t *testing.T) {
	customErr := errors.New("temporary error")

	start := time.Now()
	err := ExecWithAttempts(3, 1, mockFunc(2, customErr)) // falha 2 vezes e depois sucesso
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	// Como temos 2 falhas com delay de 1s, o tempo mínimo deve ser >=2s
	if elapsed < 2*time.Second {
		t.Fatalf("expected at least 2s elapsed, got %v", elapsed)
	}
}

func TestExecWithAttempts_AllFail(t *testing.T) {
	customErr := errors.New("persistent error")

	start := time.Now()
	err := ExecWithAttempts(3, 1, mockFunc(3, customErr)) // falha em todas as tentativas
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "max attempts reached") {
		t.Fatalf("unexpected error message: %v", err)
	}

	if !errors.Is(err, customErr) {
		t.Fatalf("expected wrapped error %v, got %v", customErr, err)
	}

	// Deve esperar 3 tentativas com 1s entre elas → no mínimo 3s (na prática ~2s porque a última não espera)
	if elapsed < 2*time.Second {
		t.Fatalf("expected at least ~2s elapsed, got %v", elapsed)
	}
}

func TestExecWithAttempts_ZeroAttempts(t *testing.T) {
	err := ExecWithAttempts(0, 1, func() error {
		return errors.New("should not be called")
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "max attempts reached") {
		t.Fatalf("unexpected error message: %v", err)
	}
}

func TestExecWithAttempts_ZeroDelay(t *testing.T) {
	customErr := errors.New("fail")
	start := time.Now()
	err := ExecWithAttempts(3, 0, mockFunc(3, customErr))
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Como delay é 0, tempo de execução deve ser <1s
	if elapsed > time.Second {
		t.Fatalf("expected elapsed <1s, got %v", elapsed)
	}
}
