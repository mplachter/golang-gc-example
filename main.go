package main

import (
	"fmt"
	"sync"

	"github.com/pkg/profile"
)

const (
	runCounts       = 10
	randomNumberMax = 4
	randomNumberMin = 1
)

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	for i := 1; i <= 5; i++ {
		DoSomeWork()
	}

}

// DoSomeWork random function that does some math work and things
func DoSomeWork() {
	nm := new(NumberStorage)

	var wg sync.WaitGroup

	for c := 1; c < runCounts; {

		for i := 1; i <= 5; i++ {
			wg.Add(1)
			go nm.worker(&wg, i, c)
		}

		wg.Wait()
		c = c + 5
		// fmt.Printf("Work Done %d/%d\n", c, runCounts)
	}

	highInts := nm.HighestOccuringNumbers()

	fmt.Printf("Highest Occuring Ints: %d\n", highInts)
}
