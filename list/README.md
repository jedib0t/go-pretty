# List
[![Go Reference](https://pkg.go.dev/badge/github.com/jedib0t/go-pretty/v6/list.svg)](https://pkg.go.dev/github.com/jedib0t/go-pretty/v6/list)

Pretty-print lists with multiple levels/indents into ASCII/Unicode strings.

## Sample List Output

```
 ■ Game Of Thrones
   ■ Winter
   ■ Is
   ■ Coming
     ■ This
     ■ Is
     ■ Known
 ■ The Dark Tower
   ■ The Gunslinger
```

A demonstration of all the capabilities can be found here:
[../cmd/demo-list](../cmd/demo-list)

## Features

### Core List Building

  - Append items one-by-one or as a group
  - Support items with multiple lines (newlines preserved)
  - Automatic handling of tabs (converted to spaces)
  - Reset list to initial state for reuse

### Indentation Control

  - Indent following items to create nested levels
  - UnIndent to move back to previous level
  - UnIndentAll to return to root level
  - Smart indentation prevents invalid nesting

### Rendering Formats

  - **ASCII/Unicode List** - Human-readable pretty format with customizable bullets and connectors
  - **HTML List** - Generate nested `<ul>` and `<li>` elements with custom CSS classes
    - Automatic CSS class numbering for nested levels (e.g., `class-1`, `class-2`)
    - HTML entity escaping for safe rendering
    - Multi-line support with `<br/>` tags
  - **Markdown List** - Generate markdown-compatible list format
    - Automatically uses markdown-appropriate styling

### Customization & Styling

  - Completely customizable styles
    - Many ready-to-use styles: [style.go](style.go)
      - `StyleDefault` - Simple asterisk bullets (*)
      - `StyleBulletCircle` - Circle bullets (●)
      - `StyleBulletFlower` - Flower bullets (✽)
      - `StyleBulletSquare` - Square bullets (■)
      - `StyleBulletStar` - Star bullets (★)
      - `StyleBulletTriangle` - Triangle bullets (▶)
      - `StyleConnectedBold` - Bold box-drawing connectors (┏━, ┣━, ┗━)
      - `StyleConnectedDouble` - Double box-drawing connectors (╔═, ╠═, ╚═)
      - `StyleConnectedLight` - Light box-drawing connectors (┌─, ├─, └─)
      - `StyleConnectedRounded` - Rounded box-drawing connectors (╭─, ├─, ╰─)
      - `StyleMarkdown` - Markdown-compatible styling
    - Customize bullet characters for different positions (top, first, middle, bottom, single)
    - Customize vertical connectors between levels
    - Set line prefix for all lines
    - Apply text formatting (colors, styles) using `text.Format`
  - HTML CSS class customization for styled HTML output

### Output Control

  - Mirror output to an io.Writer object (like os.StdOut) while rendering
  - Get rendered output as string for further processing
  - Length() method to get the number of items in the list
