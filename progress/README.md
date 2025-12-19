# Progress
[![Go Reference](https://pkg.go.dev/badge/github.com/jedib0t/go-pretty/v6/progress.svg)](https://pkg.go.dev/github.com/jedib0t/go-pretty/v6/progress)

Track the Progress of one or more Tasks (like downloading multiple files in
parallel).

## Sample Progress Tracking

<img src="images/demo.gif" width="640px"/>

A demonstration of all the capabilities can be found here:
[../cmd/demo-progress](../cmd/demo-progress)

## Features

### Core Tracking

  - Track one or more Tasks at the same time
  - Support for both determinate (with Total) and indeterminate (without Total)
    trackers
  - Error tracking with `MarkAsErrored()` for failed tasks
  - ETA (Estimated Time of Arrival) calculation for each tracker
  - Speed calculation (items/bytes per second) with customizable formatters
  - Overall progress tracker that aggregates progress across all trackers

### Tracker Management

  - Dynamically add one or more Task Trackers while `Render()` is in progress
  - Sort trackers by Index (ascending/descending), Message, Percentage, or Value
    - `SortByIndex` / `SortByIndexDsc` - Sort by explicit Index field, maintaining
      order regardless of completion status (done and active trackers are merged
      and sorted together)
    - For other sorting methods, done and active trackers are sorted separately,
      with done trackers always rendered before active trackers
  - Tracker options
    - `AutoStopDisabled` - Prevent auto-completion when value exceeds total
    - `DeferStart` - Delay tracker start until manually triggered
    - `Index` - Explicit ordering value for trackers (used with `SortByIndex`)
    - `RemoveOnCompletion` - Hide tracker when done instead of showing completion

### Display & Rendering

  - Choose to have the Writer auto-stop the Render when no more Trackers are
    in queue, or manually stop using `Stop()`
  - Redirect output to an io.Writer object (like os.StdOut)
  - Pinned messages that stay visible above all trackers
  - Log messages during rendering that appear temporarily
  - Automatic terminal width detection to prevent line wrapping
  - Flexible tracker positioning (left or right of message)
  - Configurable update frequency for smooth rendering

### Customization & Styling

  - Completely customizable styles
    - Many ready-to-use styles: [style.go](style.go)
      - `StyleDefault` - ASCII characters
      - `StyleBlocks` - UNICODE Block Drawing characters
      - `StyleCircle` - UNICODE Circle runes
      - `StyleRhombus` - UNICODE Rhombus runes
    - Colorize various parts of the Tracker using `StyleColors`
    - Customize how Trackers get rendered using `StyleOptions`
    - Control visibility of components (ETA, Speed, Time, Value, etc.)
    - Custom renderers for determinate and indeterminate progress bars
  - Multiple indeterminate indicator animations
    - Moving back and forth
    - Moving left to right
    - Moving right to left
    - Dominoes effect
    - Pac-Man chomping animation
    - Colored variants

### Units & Formatting

  - Built-in unit formatters
    - `UnitsDefault` - Regular numbers (K, M, B, T, Q notation)
    - `UnitsBytes` - Storage units (B, KB, MB, GB, TB, PB)
    - `UnitsCurrencyDollar` - Dollar amounts ($x.yzK, etc.)
    - `UnitsCurrencyEuro` - Euro amounts (₠x.yzK, etc.)
    - `UnitsCurrencyPound` - Pound amounts (£x.yzK, etc.)
