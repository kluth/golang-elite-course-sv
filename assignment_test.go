package assignment

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestRetry(t *testing.T) {
	c := 0
	err := Retry(func() error {
		c++
		if c < 3 { return errors.New("fail") }
		return nil
	}, 5)
	if err != nil || c != 3 {
		t.Fatalf("failed")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout d8016bc88a9f26de")
	}
	os.Exit(code)
}
