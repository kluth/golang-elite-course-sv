package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestUser(t *testing.T) {
	u := NewUser("123")
	if u.ID() != "123" {
		t.Fatalf("failed")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout 16f8c38be20d795c")
	}
	os.Exit(code)
}
