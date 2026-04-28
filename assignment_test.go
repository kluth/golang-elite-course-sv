package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestPatterns(t *testing.T) {
	if (Facade{}).Unified() != "12" { t.Fatalf("failed") }
	p := Proxy{Role: "user"}
	if p.Get() != "denied" { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout 417ab8490c82b827")
	}
	os.Exit(code)
}
