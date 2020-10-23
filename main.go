package main

import (
	"github.com/pkg/profile"
)

const (
	runTimes = 10
)

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	for i := 1; i <= runTimes; i++ {
		DoSomeWork()
	}

}
