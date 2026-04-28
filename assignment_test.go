package assignment

import (
	"fmt"
	"os"
	"testing"
)

type mockVis struct{}
func (m mockVis) VisitCircle() string { return "v" }

func TestPatterns(t *testing.T) {
	if (Circle{}).Accept(mockVis{}) != "v" { t.Fatalf("failed") }
	if GetTreeType("a") != GetTreeType("a") { t.Fatalf("failed") }
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout COURSE_COMPLETE")
	}
	os.Exit(code)
}
