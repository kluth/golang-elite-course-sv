package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerics(t *testing.T) {
	s := NewSet[string]()
	s.Add("go")
	if !s["go"] {
		t.Fatalf("failed")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout cdfc9f4784354c43")
	}
	os.Exit(code)
}
