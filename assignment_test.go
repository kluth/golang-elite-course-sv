package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestPatterns(t *testing.T) {
	fld := &Folder{}; fld.Add(File{10}); fld.Add(File{20})
	if fld.Size() != 30 { t.Fatalf("failed") }
	dec := LoggingDecorator{ConcreteService{}}
	if dec.Do() != "log-do" { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout ec38b6e5aed6137c")
	}
	os.Exit(code)
}
