package logger_go_test

import (
	"sync"
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

func TestRace(t *testing.T) {
	dest := "./dest"
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			m := logger_go.New(dest, dest, dest, dest)
			if m == nil {
				t.Error("no se pudo crear el logger")
			}
			m.Info.Println("estoy desde gorutinas: ", i)
		}(wg, i)
	}
	wg.Wait()
}
