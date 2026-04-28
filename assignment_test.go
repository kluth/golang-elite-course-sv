package assignment

import (
	"fmt"
	"os"
	"testing"
)

func TestWorker(t *testing.T) {
	j := make(chan int, 2)
	r := make(chan int, 2)
	go WorkerPool(1, j, r)
	j <- 1; j <- 2; close(j)
	<-r; <-r
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		fmt.Println("\nUPPDRAG SLUTFÖRT. Fortsätt till:")
		fmt.Println("git checkout a7a6b9495b789fcf")
	}
	os.Exit(code)
}
