package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestInterface(t *testing.T) {
	var n Notifier = &EmailSender{}
	if n.Notify() != "email" {
		t.Fatalf("failed")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout 1e4e2e5da68726b4")
	}
	os.Exit(code)
}
