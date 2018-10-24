package list

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Render(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem(testItem1)
	lw.Indent()
	lw.AppendItems(testItems2)
	lw.Indent()
	lw.AppendItems(testItems3)
	lw.UnIndent()
	lw.UnIndent()
	lw.AppendItem(testItem4)
	lw.Indent()
	lw.AppendItem(testItem5)
	lw.SetStyle(styleTest)

	expectedOut := `t Game Of Thrones
|f Winter
|m Is
|b Coming
| f This
| m Is
| b Known
b The Dark Tower
 b The Gunslinger`
	assert.Equal(t, expectedOut, lw.Render())
}

func TestList_Render_Complex(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem("The Houses of Westeros")
	lw.Indent()
	lw.AppendItem("The Starks of Winterfell")
	lw.Indent()
	lw.AppendItem("Eddard Stark")
	lw.Indent()
	lw.AppendItems([]interface{}{"Robb Stark", "Sansa Stark", "Arya Stark", "Bran Stark", "Rickon Stark"})
	lw.UnIndent()
	lw.AppendItems([]interface{}{"Lyanna Stark", "Benjen Stark"})
	lw.UnIndent()
	lw.AppendItem("The Targaryens of Dragonstone")
	lw.Indent()
	lw.AppendItem("Aerys Targaryen")
	lw.Indent()
	lw.AppendItems([]interface{}{"Rhaegar Targaryen", "Viserys Targaryen", "Daenerys Targaryen"})
	lw.UnIndent()
	lw.UnIndent()
	lw.AppendItem("The Lannisters of Lannisport")
	lw.Indent()
	lw.AppendItem("Tywin Lannister")
	lw.Indent()
	lw.AppendItems([]interface{}{"Cersei Lannister", "Jaime Lannister", "Tyrion Lannister"})

	styles := map[Style]string{
		StyleBulletCircle:     "● The Houses of Westeros\n  ● The Starks of Winterfell\n    ● Eddard Stark\n      ● Robb Stark\n      ● Sansa Stark\n      ● Arya Stark\n      ● Bran Stark\n      ● Rickon Stark\n    ● Lyanna Stark\n    ● Benjen Stark\n  ● The Targaryens of Dragonstone\n    ● Aerys Targaryen\n      ● Rhaegar Targaryen\n      ● Viserys Targaryen\n      ● Daenerys Targaryen\n  ● The Lannisters of Lannisport\n    ● Tywin Lannister\n      ● Cersei Lannister\n      ● Jaime Lannister\n      ● Tyrion Lannister",
		StyleBulletFlower:     "✽ The Houses of Westeros\n  ✽ The Starks of Winterfell\n    ✽ Eddard Stark\n      ✽ Robb Stark\n      ✽ Sansa Stark\n      ✽ Arya Stark\n      ✽ Bran Stark\n      ✽ Rickon Stark\n    ✽ Lyanna Stark\n    ✽ Benjen Stark\n  ✽ The Targaryens of Dragonstone\n    ✽ Aerys Targaryen\n      ✽ Rhaegar Targaryen\n      ✽ Viserys Targaryen\n      ✽ Daenerys Targaryen\n  ✽ The Lannisters of Lannisport\n    ✽ Tywin Lannister\n      ✽ Cersei Lannister\n      ✽ Jaime Lannister\n      ✽ Tyrion Lannister",
		StyleBulletSquare:     "■ The Houses of Westeros\n  ■ The Starks of Winterfell\n    ■ Eddard Stark\n      ■ Robb Stark\n      ■ Sansa Stark\n      ■ Arya Stark\n      ■ Bran Stark\n      ■ Rickon Stark\n    ■ Lyanna Stark\n    ■ Benjen Stark\n  ■ The Targaryens of Dragonstone\n    ■ Aerys Targaryen\n      ■ Rhaegar Targaryen\n      ■ Viserys Targaryen\n      ■ Daenerys Targaryen\n  ■ The Lannisters of Lannisport\n    ■ Tywin Lannister\n      ■ Cersei Lannister\n      ■ Jaime Lannister\n      ■ Tyrion Lannister",
		StyleBulletStar:       "★ The Houses of Westeros\n  ★ The Starks of Winterfell\n    ★ Eddard Stark\n      ★ Robb Stark\n      ★ Sansa Stark\n      ★ Arya Stark\n      ★ Bran Stark\n      ★ Rickon Stark\n    ★ Lyanna Stark\n    ★ Benjen Stark\n  ★ The Targaryens of Dragonstone\n    ★ Aerys Targaryen\n      ★ Rhaegar Targaryen\n      ★ Viserys Targaryen\n      ★ Daenerys Targaryen\n  ★ The Lannisters of Lannisport\n    ★ Tywin Lannister\n      ★ Cersei Lannister\n      ★ Jaime Lannister\n      ★ Tyrion Lannister",
		StyleBulletTriangle:   "▶ The Houses of Westeros\n  ▶ The Starks of Winterfell\n    ▶ Eddard Stark\n      ▶ Robb Stark\n      ▶ Sansa Stark\n      ▶ Arya Stark\n      ▶ Bran Stark\n      ▶ Rickon Stark\n    ▶ Lyanna Stark\n    ▶ Benjen Stark\n  ▶ The Targaryens of Dragonstone\n    ▶ Aerys Targaryen\n      ▶ Rhaegar Targaryen\n      ▶ Viserys Targaryen\n      ▶ Daenerys Targaryen\n  ▶ The Lannisters of Lannisport\n    ▶ Tywin Lannister\n      ▶ Cersei Lannister\n      ▶ Jaime Lannister\n      ▶ Tyrion Lannister",
		StyleConnectedBold:    "━━ The Houses of Westeros\n   ┣━ The Starks of Winterfell\n   ┃  ┣━ Eddard Stark\n   ┃  ┃  ┣━ Robb Stark\n   ┃  ┃  ┣━ Sansa Stark\n   ┃  ┃  ┣━ Arya Stark\n   ┃  ┃  ┣━ Bran Stark\n   ┃  ┃  ┗━ Rickon Stark\n   ┃  ┣━ Lyanna Stark\n   ┃  ┗━ Benjen Stark\n   ┣━ The Targaryens of Dragonstone\n   ┃  ┗━ Aerys Targaryen\n   ┃     ┣━ Rhaegar Targaryen\n   ┃     ┣━ Viserys Targaryen\n   ┃     ┗━ Daenerys Targaryen\n   ┗━ The Lannisters of Lannisport\n      ┗━ Tywin Lannister\n         ┣━ Cersei Lannister\n         ┣━ Jaime Lannister\n         ┗━ Tyrion Lannister",
		StyleConnectedDouble:  "══ The Houses of Westeros\n   ╠═ The Starks of Winterfell\n   ║  ╠═ Eddard Stark\n   ║  ║  ╠═ Robb Stark\n   ║  ║  ╠═ Sansa Stark\n   ║  ║  ╠═ Arya Stark\n   ║  ║  ╠═ Bran Stark\n   ║  ║  ╚═ Rickon Stark\n   ║  ╠═ Lyanna Stark\n   ║  ╚═ Benjen Stark\n   ╠═ The Targaryens of Dragonstone\n   ║  ╚═ Aerys Targaryen\n   ║     ╠═ Rhaegar Targaryen\n   ║     ╠═ Viserys Targaryen\n   ║     ╚═ Daenerys Targaryen\n   ╚═ The Lannisters of Lannisport\n      ╚═ Tywin Lannister\n         ╠═ Cersei Lannister\n         ╠═ Jaime Lannister\n         ╚═ Tyrion Lannister",
		StyleConnectedLight:   "── The Houses of Westeros\n   ├─ The Starks of Winterfell\n   │  ├─ Eddard Stark\n   │  │  ├─ Robb Stark\n   │  │  ├─ Sansa Stark\n   │  │  ├─ Arya Stark\n   │  │  ├─ Bran Stark\n   │  │  └─ Rickon Stark\n   │  ├─ Lyanna Stark\n   │  └─ Benjen Stark\n   ├─ The Targaryens of Dragonstone\n   │  └─ Aerys Targaryen\n   │     ├─ Rhaegar Targaryen\n   │     ├─ Viserys Targaryen\n   │     └─ Daenerys Targaryen\n   └─ The Lannisters of Lannisport\n      └─ Tywin Lannister\n         ├─ Cersei Lannister\n         ├─ Jaime Lannister\n         └─ Tyrion Lannister",
		StyleConnectedRounded: "── The Houses of Westeros\n   ├─ The Starks of Winterfell\n   │  ├─ Eddard Stark\n   │  │  ├─ Robb Stark\n   │  │  ├─ Sansa Stark\n   │  │  ├─ Arya Stark\n   │  │  ├─ Bran Stark\n   │  │  ╰─ Rickon Stark\n   │  ├─ Lyanna Stark\n   │  ╰─ Benjen Stark\n   ├─ The Targaryens of Dragonstone\n   │  ╰─ Aerys Targaryen\n   │     ├─ Rhaegar Targaryen\n   │     ├─ Viserys Targaryen\n   │     ╰─ Daenerys Targaryen\n   ╰─ The Lannisters of Lannisport\n      ╰─ Tywin Lannister\n         ├─ Cersei Lannister\n         ├─ Jaime Lannister\n         ╰─ Tyrion Lannister",
		StyleDefault:          "* The Houses of Westeros\n  * The Starks of Winterfell\n    * Eddard Stark\n      * Robb Stark\n      * Sansa Stark\n      * Arya Stark\n      * Bran Stark\n      * Rickon Stark\n    * Lyanna Stark\n    * Benjen Stark\n  * The Targaryens of Dragonstone\n    * Aerys Targaryen\n      * Rhaegar Targaryen\n      * Viserys Targaryen\n      * Daenerys Targaryen\n  * The Lannisters of Lannisport\n    * Tywin Lannister\n      * Cersei Lannister\n      * Jaime Lannister\n      * Tyrion Lannister",
		StyleMarkdown:         "  * The Houses of Westeros\n    * The Starks of Winterfell\n      * Eddard Stark\n        * Robb Stark\n        * Sansa Stark\n        * Arya Stark\n        * Bran Stark\n        * Rickon Stark\n      * Lyanna Stark\n      * Benjen Stark\n    * The Targaryens of Dragonstone\n      * Aerys Targaryen\n        * Rhaegar Targaryen\n        * Viserys Targaryen\n        * Daenerys Targaryen\n    * The Lannisters of Lannisport\n      * Tywin Lannister\n        * Cersei Lannister\n        * Jaime Lannister\n        * Tyrion Lannister",
		styleTest:             "s The Houses of Westeros\n f The Starks of Winterfell\n |f Eddard Stark\n ||f Robb Stark\n ||m Sansa Stark\n ||m Arya Stark\n ||m Bran Stark\n ||b Rickon Stark\n |m Lyanna Stark\n |b Benjen Stark\n m The Targaryens of Dragonstone\n |b Aerys Targaryen\n | f Rhaegar Targaryen\n | m Viserys Targaryen\n | b Daenerys Targaryen\n b The Lannisters of Lannisport\n  b Tywin Lannister\n   f Cersei Lannister\n   m Jaime Lannister\n   b Tyrion Lannister",
	}
	var mismatches []string
	for style, expectedOut := range styles {
		lw.SetStyle(style)
		out := lw.Render()
		assert.Equal(t, expectedOut, out)
		if expectedOut != out {
			mismatches = append(mismatches, fmt.Sprintf("%s: %#v,", style.Name, out))
			fmt.Printf("// %s renders a List like below:\n", style.Name)
			for _, line := range strings.Split(out, "\n") {
				fmt.Printf("//  %s\n", line)
			}
			fmt.Println()
		}
	}
	sort.Strings(mismatches)
	for _, mismatch := range mismatches {
		fmt.Println(mismatch)
	}
}

func TestList_Render_Connected(t *testing.T) {
	lw := NewWriter()
	lw.SetStyle(StyleConnectedLight)
	assert.Empty(t, lw.Render())

	lw.AppendItem(testItem1)
	expectedOut := "── Game Of Thrones"
	assert.Equal(t, expectedOut, lw.Render())

	lw.AppendItem(testItem1)
	expectedOut = `┌─ Game Of Thrones
└─ Game Of Thrones`
	assert.Equal(t, expectedOut, lw.Render())

	lw.AppendItem(testItem1)
	expectedOut = `┌─ Game Of Thrones
├─ Game Of Thrones
└─ Game Of Thrones`
	assert.Equal(t, expectedOut, lw.Render())

	lw.Indent()
	lw.AppendItem(testItem1)
	expectedOut = `┌─ Game Of Thrones
├─ Game Of Thrones
└─ Game Of Thrones
   └─ Game Of Thrones`
	assert.Equal(t, expectedOut, lw.Render())

	lw.AppendItem(testItem1)
	expectedOut = `┌─ Game Of Thrones
├─ Game Of Thrones
└─ Game Of Thrones
   ├─ Game Of Thrones
   └─ Game Of Thrones`
	assert.Equal(t, expectedOut, lw.Render())

	lw.Indent()
	lw.AppendItem(testItem1)
	expectedOut = `┌─ Game Of Thrones
├─ Game Of Thrones
└─ Game Of Thrones
   ├─ Game Of Thrones
   └─ Game Of Thrones
      └─ Game Of Thrones`
	assert.Equal(t, expectedOut, lw.Render())

	lw.UnIndent()
	lw.AppendItem(testItem1)
	expectedOut = `┌─ Game Of Thrones
├─ Game Of Thrones
└─ Game Of Thrones
   ├─ Game Of Thrones
   ├─ Game Of Thrones
   │  └─ Game Of Thrones
   └─ Game Of Thrones`
	assert.Equal(t, expectedOut, lw.Render())

	lw.UnIndent()
	lw.AppendItem(testItem1)
	expectedOut = `┌─ Game Of Thrones
├─ Game Of Thrones
├─ Game Of Thrones
│  ├─ Game Of Thrones
│  ├─ Game Of Thrones
│  │  └─ Game Of Thrones
│  └─ Game Of Thrones
└─ Game Of Thrones`
	assert.Equal(t, expectedOut, lw.Render())
}

func TestList_Render_MultiLine(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem(testItem1ML)
	lw.Indent()
	lw.AppendItems(testItems2ML)
	lw.Indent()
	lw.AppendItems(testItems3ML)
	lw.UnIndent()
	lw.UnIndent()
	lw.AppendItem(testItem4ML)
	lw.Indent()
	lw.AppendItem(testItem5)

	expectedOut := `* Game Of Thrones
      // George. R. R. Martin
  * Winter
    Is
    Coming
  * Is
  * Coming
    * This
      Is
      Known
    * Is
    * Known
* The Dark Tower
      // Stephen King
  * The Gunslinger`
	assert.Equal(t, expectedOut, lw.Render())

	expectedOutRounded := `╭─ Game Of Thrones
│      // George. R. R. Martin
│  ├─ Winter
│  │  Is
│  │  Coming
│  ├─ Is
│  ╰─ Coming
│     ├─ This
│     │  Is
│     │  Known
│     ├─ Is
│     ╰─ Known
╰─ The Dark Tower
       // Stephen King
   ╰─ The Gunslinger`
	lw.SetStyle(StyleConnectedRounded)
	assert.Equal(t, expectedOutRounded, lw.Render())

	expectedOutHTML := `<ul class="go-pretty-table">
  <li>Game Of Thrones<br/>    // George. R. R. Martin</li>
  <ul class="go-pretty-table-1">
    <li>Winter<br/>Is<br/>Coming</li>
    <li>Is</li>
    <li>Coming</li>
    <ul class="go-pretty-table-2">
      <li>This<br/>Is<br/>Known</li>
      <li>Is</li>
      <li>Known</li>
    </ul>
  </ul>
  <li>The Dark Tower<br/>    // Stephen King</li>
  <ul class="go-pretty-table-1">
    <li>The Gunslinger</li>
  </ul>
</ul>`
	assert.Equal(t, expectedOutHTML, lw.RenderHTML())

	expectedOutMarkdown := `  * Game Of Thrones<br/>    // George. R. R. Martin
    * Winter<br/>Is<br/>Coming
    * Is
    * Coming
      * This<br/>Is<br/>Known
      * Is
      * Known
  * The Dark Tower<br/>    // Stephen King
    * The Gunslinger`
	assert.Equal(t, expectedOutMarkdown, lw.RenderMarkdown())
}

func TestList_Render_Styles(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem(testItem1)
	lw.Indent()
	lw.AppendItems(testItems2)
	lw.Indent()
	lw.AppendItems(testItems3)
	lw.UnIndent()
	lw.UnIndent()
	lw.AppendItem(testItem4)
	lw.Indent()
	lw.AppendItem(testItem5)

	styles := map[Style]string{
		StyleBulletCircle:     "● Game Of Thrones\n  ● Winter\n  ● Is\n  ● Coming\n    ● This\n    ● Is\n    ● Known\n● The Dark Tower\n  ● The Gunslinger",
		StyleBulletFlower:     "✽ Game Of Thrones\n  ✽ Winter\n  ✽ Is\n  ✽ Coming\n    ✽ This\n    ✽ Is\n    ✽ Known\n✽ The Dark Tower\n  ✽ The Gunslinger",
		StyleBulletSquare:     "■ Game Of Thrones\n  ■ Winter\n  ■ Is\n  ■ Coming\n    ■ This\n    ■ Is\n    ■ Known\n■ The Dark Tower\n  ■ The Gunslinger",
		StyleBulletStar:       "★ Game Of Thrones\n  ★ Winter\n  ★ Is\n  ★ Coming\n    ★ This\n    ★ Is\n    ★ Known\n★ The Dark Tower\n  ★ The Gunslinger",
		StyleBulletTriangle:   "▶ Game Of Thrones\n  ▶ Winter\n  ▶ Is\n  ▶ Coming\n    ▶ This\n    ▶ Is\n    ▶ Known\n▶ The Dark Tower\n  ▶ The Gunslinger",
		StyleConnectedBold:    "┏━ Game Of Thrones\n┃  ┣━ Winter\n┃  ┣━ Is\n┃  ┗━ Coming\n┃     ┣━ This\n┃     ┣━ Is\n┃     ┗━ Known\n┗━ The Dark Tower\n   ┗━ The Gunslinger",
		StyleConnectedDouble:  "╔═ Game Of Thrones\n║  ╠═ Winter\n║  ╠═ Is\n║  ╚═ Coming\n║     ╠═ This\n║     ╠═ Is\n║     ╚═ Known\n╚═ The Dark Tower\n   ╚═ The Gunslinger",
		StyleConnectedLight:   "┌─ Game Of Thrones\n│  ├─ Winter\n│  ├─ Is\n│  └─ Coming\n│     ├─ This\n│     ├─ Is\n│     └─ Known\n└─ The Dark Tower\n   └─ The Gunslinger",
		StyleConnectedRounded: "╭─ Game Of Thrones\n│  ├─ Winter\n│  ├─ Is\n│  ╰─ Coming\n│     ├─ This\n│     ├─ Is\n│     ╰─ Known\n╰─ The Dark Tower\n   ╰─ The Gunslinger",
		StyleDefault:          "* Game Of Thrones\n  * Winter\n  * Is\n  * Coming\n    * This\n    * Is\n    * Known\n* The Dark Tower\n  * The Gunslinger",
		StyleMarkdown:         "  * Game Of Thrones\n    * Winter\n    * Is\n    * Coming\n      * This\n      * Is\n      * Known\n  * The Dark Tower\n    * The Gunslinger",
		styleTest:             "t Game Of Thrones\n|f Winter\n|m Is\n|b Coming\n| f This\n| m Is\n| b Known\nb The Dark Tower\n b The Gunslinger",
	}
	var mismatches []string
	for style, expectedOut := range styles {
		lw.SetStyle(style)
		out := lw.Render()
		assert.Equal(t, expectedOut, out)
		if expectedOut != out {
			mismatches = append(mismatches, fmt.Sprintf("%s: %#v,", style.Name, out))
			fmt.Printf("// %s renders a List like below:\n", style.Name)
			for _, line := range strings.Split(out, "\n") {
				fmt.Printf("//  %s\n", line)
			}
			fmt.Println()
		}
	}
	sort.Strings(mismatches)
	for _, mismatch := range mismatches {
		fmt.Println(mismatch)
	}
}
