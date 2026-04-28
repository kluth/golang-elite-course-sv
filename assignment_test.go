package assignment

import (
	"fmt"
	"os"
	"testing"
)

type mockWork struct { c int }
func (m *mockWork) Step1() { m.c++ }; func (m *mockWork) Step2() { m.c++ }

func TestPatterns(t *testing.T) {
	mw := &mockWork{}; ExecuteWork(mw)
	if mw.c != 2 { t.Fatalf("failed") }
	g := &Group{items: []string{"a"}}
	if !g.HasNext() || g.Next() != "a" { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout 746196646fcecbe8")
	}
	os.Exit(code)
}
