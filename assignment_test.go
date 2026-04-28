package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestSlice(t *testing.T) {
	orig := []int{1, 2, 3}
	mod := CloneAndModify(orig, 0, 99)
	if orig[0] == 99 || mod[0] != 99 {
		t.Fatalf("failed")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout bc4b2be876538b0c")
	}
	os.Exit(code)
}
