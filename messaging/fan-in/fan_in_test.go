package fan_in

import (
	"testing"
	"time"
	"fmt"
)

func TestFanin(t *testing.T){
	done := make(chan bool)
	defer close(done)

	start := time.Now()

	items := PrepareItems(done)

	workers := make([]<-chan int, 8)
	for i := 0; i<8; i++ {
		workers[i] = PackItems(done, items, i)
	}

	numPackages := 0
	for range merge(done, workers...) {
		numPackages++
	}

	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
