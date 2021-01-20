# Progress
[![Go Reference](https://pkg.go.dev/badge/github.com/jedib0t/go-pretty/v6/progress.svg)](https://pkg.go.dev/github.com/jedib0t/go-pretty/v6/progress)

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

A demonstration of all the capabilities can be found here:
[../cmd/demo-progress](../cmd/demo-progress)

## Sample Progress Tracking

<img src="images/demo.gif" width="640px"/>

# TODO

  - Optimize CPU and Memory Usage
