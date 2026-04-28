package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1}
	if _, ok := LookupSafe(m, "b"); ok {
		t.Fatalf("failed")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout 3f18728a4d445c95")
	}
	os.Exit(code)
}
