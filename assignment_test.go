package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestPatterns(t *testing.T) {
	n := Navigator{Walk{}}
	if n.Exec() != "walk" { t.Fatalf("failed") }
	o := Order{Pending{}}
	if o.Process() != "pending" { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout 4e5ffd7df3dcc20b")
	}
	os.Exit(code)
}
