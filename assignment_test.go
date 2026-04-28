package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestPatterns(t *testing.T) {
	var s ModernService = LegacyAdapter{LegacySystem{}}
	if s.Call() != "legacy" { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout 8cc7f0dcddba46b5")
	}
	os.Exit(code)
}
