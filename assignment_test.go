package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestOptions(t *testing.T) {
	s := NewServer(WithPort(8080))
	if s.Port != 8080 {
		t.Fatalf("failed")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout 89bd806ab79e81c5")
	}
	os.Exit(code)
}
