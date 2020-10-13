package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/pkg/profile"
)

var (
	tracker1  = progress.Tracker{Message: "Calculating Total   # 1", Total: 1000, Units: progress.UnitsDefault}
	tracker2  = progress.Tracker{Message: "Downloading File    # 2", Total: 1000, Units: progress.UnitsBytes}
	tracker3  = progress.Tracker{Message: "Transferring Amount # 3", Total: 1000, Units: progress.UnitsCurrencyDollar}
	profilers = []func(*profile.Profile){
		profile.CPUProfile,
		profile.MemProfileRate(512),
	}
)

func profileRender(profiler func(profile2 *profile.Profile), n int) {
	defer profile.Start(profiler, profile.ProfilePath("./")).Stop()

	trackSomething := func(pw progress.Writer, tracker *progress.Tracker) {
		tracker.Reset()
		pw.AppendTracker(tracker)
		time.Sleep(time.Millisecond * 100)
		tracker.Increment(tracker.Total / 2)
		time.Sleep(time.Millisecond * 100)
		tracker.Increment(tracker.Total / 2)
	}

	for i := 0; i < n; i++ {
		pw := progress.NewWriter()
		pw.SetAutoStop(true)
		pw.SetOutputWriter(ioutil.Discard)
		go trackSomething(pw, &tracker1)
		go trackSomething(pw, &tracker2)
		go trackSomething(pw, &tracker3)
		time.Sleep(time.Millisecond * 50)
		pw.Render()
	}
}

func main() {
	numRenders := 5
	if len(os.Args) > 1 {
		var err error
		numRenders, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Invalid Argument: '%s'\n", os.Args[2])
			os.Exit(1)
		}
	}

	for _, profiler := range profilers {
		profileRender(profiler, numRenders)
	}
}
