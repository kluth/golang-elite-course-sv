package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestSafeAdd(t *testing.T) {
	tests := []struct{ a, b, want int32; err bool }{
		{10, 20, 30, false},
		{2147483647, 1, 0, true},
	}
	for _, tt := range tests {
		got, err := SafeAdd(tt.a, tt.b)
		if (err != nil) != tt.err || (!tt.err && got != tt.want) {
			t.Errorf("failed")
		}
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout 1cc747eaceff9ee6")
	}
	os.Exit(code)
}
