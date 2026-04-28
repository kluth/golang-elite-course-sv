package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestPatterns(t *testing.T) {
	if ProductFactory("any").GetName() != "concrete" { t.Fatalf("failed") }
	b := &CarBuilder{}; c := b.SetColor("red").Build()
	if c.Color != "red" { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout d2412a134d7b3f02")
	}
	os.Exit(code)
}
