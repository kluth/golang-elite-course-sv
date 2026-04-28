package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestPatterns(t *testing.T) {
	if GetInstance() != GetInstance() { t.Fatalf("failed") }
	u1 := &User{"A"}; u2 := u1.Clone()
	if u1 == u2 || u2.Name != "A" { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout bfb4238b82b207a5")
	}
	os.Exit(code)
}
