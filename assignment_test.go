package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestPatterns(t *testing.T) {
	h := &BaseHandler{}
	if h.Handle("any") { t.Fatalf("failed") }
	if (LightOn{}).Execute() != "on" { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout b241a02ef96deebe")
	}
	os.Exit(code)
}
