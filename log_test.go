package logger_go_test

import (
	"testing"

	"github.com/alexyslozada/logger-go"
)

func TestNew(t *testing.T) {
	dest := "./dest"
	m := logger_go.New(dest, dest, dest, dest)
	if m == nil {
		t.Error("No se pudo crear los archivos")
	}
}
