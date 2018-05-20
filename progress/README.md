# Progress
[![GoDoc](https://godoc.org/github.com/jedib0t/go-pretty/progress?status.svg)](https://godoc.org/github.com/jedib0t/go-pretty/progress)

Track the Progress of one or more Tasks (like downloading multiple files in
parallel).

  - Track one or more Tasks at the same time
  - Dynamically add one or more Task Trackers while `Render()` is in progress
  - Choose to have the Writer auto-stop the Render when no more Trackers are
    in queue, or manually stop using `Stop()`
  - Redirect output to an io.Writer object (like os.StdOut)
  - Completely customizable styles
    - Many ready-to-use styles: [style.go](style.go)
    - Colorize various parts of the Tracker using `StyleColors`
    - Customize how Trackers get rendered using `StyleOptions`

Sample Progress Tracking:
```
Calculating Total   #  1 ... done! [3.25K in 100ms]
Calculating Total   #  2 ... done! [6.50K in 100ms]
Downloading File    #  3 ... done! [9.75KB in 100ms]
Transferring Amount #  4 ... done! [$26.00K in 200ms]
Transferring Amount #  5 ... done! [£32.50K in 201ms]
Downloading File    #  6 ... done! [58.50KB in 300ms]
Calculating Total   #  7 ... done! [91.00K in 400ms]
Transferring Amount #  8 ... 60.9% (●●●●●●●●●●●●●●◌◌◌◌◌◌◌◌◌) [$78.00K in 399.071ms]
Downloading File    #  9 ... 32.1% (●●●●●●●○◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌) [58.50KB in 298.947ms]
Transferring Amount # 10 ... 13.0% (●●○◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌◌) [£32.50K in 198.84ms]
```

Real-time playback of the demo @ asciinema.org:
[![asciicast](https://asciinema.org/a/KcPw8aoBSsYCBOj60wluhu5z3.png)](https://asciinema.org/a/KcPw8aoBSsYCBOj60wluhu5z3)

A demonstration of all the capabilities can be found here:
[../cmd/demo-progress](../cmd/demo-progress)

# TODO

  - Optimize CPU and Memory Usage
