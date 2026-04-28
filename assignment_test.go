package assignment

import (
	"fmt"
	"os"
	"testing"
)

type mockObs struct { msg string }
func (m *mockObs) Update(s string) { m.msg = s }

func TestPatterns(t *testing.T) {
	o := &mockObs{}; s := &Subject{obs: []Observer{o}}
	s.Notify("hi"); if o.msg != "hi" { t.Fatalf("failed") }
	if (ChatRoom{}).Send("a", "b") != "b:a" { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout ade37f54472542e0")
	}
	os.Exit(code)
}
