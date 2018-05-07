package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
)

func TestList_Render(t *testing.T) {
	lw := NewWriter()
	lw.AppendItem(testItem1)
	lw.Indent()
	lw.AppendItems(testItems2)
	lw.Indent()
	lw.AppendItems(testItems3)
	lw.UnIndent()
	lw.AppendItem(testItem4)
	lw.Indent()
	lw.AppendItem(testItem5)
	lw.SetStyle(styleTest)

	expectedOut := `^> Game Of Thrones
c~f> Winter
  i> Is
  i> Coming
  T~f> This
  | i> Is
  | v> Known
  i> The Dark Tower
  c~I> The Gunslinger`

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
		StyleConnectedBold:    "┏━ The Houses of Westeros\n┗━┳━ The Starks of Winterfell\n  ┣━┳━ Eddard Stark\n  ┃ ┣━┳━ Robb Stark\n  ┃ ┃ ┣━ Sansa Stark\n  ┃ ┃ ┣━ Arya Stark\n  ┃ ┃ ┣━ Bran Stark\n  ┃ ┃ ┗━ Rickon Stark\n  ┃ ┣━ Lyanna Stark\n  ┃ ┗━ Benjen Stark\n  ┣━ The Targaryens of Dragonstone\n  ┣━┳━ Aerys Targaryen\n  ┃ ┗━┳━ Rhaegar Targaryen\n  ┃   ┣━ Viserys Targaryen\n  ┃   ┗━ Daenerys Targaryen\n  ┣━ The Lannisters of Lannisport\n  ┗━┳━ Tywin Lannister\n    ┗━┳━ Cersei Lannister\n      ┣━ Jaime Lannister\n      ┗━ Tyrion Lannister",
		StyleConnectedDouble:  "╔═ The Houses of Westeros\n╚═╦═ The Starks of Winterfell\n  ╠═╦═ Eddard Stark\n  ║ ╠═╦═ Robb Stark\n  ║ ║ ╠═ Sansa Stark\n  ║ ║ ╠═ Arya Stark\n  ║ ║ ╠═ Bran Stark\n  ║ ║ ╚═ Rickon Stark\n  ║ ╠═ Lyanna Stark\n  ║ ╚═ Benjen Stark\n  ╠═ The Targaryens of Dragonstone\n  ╠═╦═ Aerys Targaryen\n  ║ ╚═╦═ Rhaegar Targaryen\n  ║   ╠═ Viserys Targaryen\n  ║   ╚═ Daenerys Targaryen\n  ╠═ The Lannisters of Lannisport\n  ╚═╦═ Tywin Lannister\n    ╚═╦═ Cersei Lannister\n      ╠═ Jaime Lannister\n      ╚═ Tyrion Lannister",
		StyleConnectedLight:   "┌─ The Houses of Westeros\n└─┬─ The Starks of Winterfell\n  ├─┬─ Eddard Stark\n  │ ├─┬─ Robb Stark\n  │ │ ├─ Sansa Stark\n  │ │ ├─ Arya Stark\n  │ │ ├─ Bran Stark\n  │ │ └─ Rickon Stark\n  │ ├─ Lyanna Stark\n  │ └─ Benjen Stark\n  ├─ The Targaryens of Dragonstone\n  ├─┬─ Aerys Targaryen\n  │ └─┬─ Rhaegar Targaryen\n  │   ├─ Viserys Targaryen\n  │   └─ Daenerys Targaryen\n  ├─ The Lannisters of Lannisport\n  └─┬─ Tywin Lannister\n    └─┬─ Cersei Lannister\n      ├─ Jaime Lannister\n      └─ Tyrion Lannister",
		StyleConnectedRounded: "╭─ The Houses of Westeros\n╰─┬─ The Starks of Winterfell\n  ├─┬─ Eddard Stark\n  │ ├─┬─ Robb Stark\n  │ │ ├─ Sansa Stark\n  │ │ ├─ Arya Stark\n  │ │ ├─ Bran Stark\n  │ │ ╰─ Rickon Stark\n  │ ├─ Lyanna Stark\n  │ ╰─ Benjen Stark\n  ├─ The Targaryens of Dragonstone\n  ├─┬─ Aerys Targaryen\n  │ ╰─┬─ Rhaegar Targaryen\n  │   ├─ Viserys Targaryen\n  │   ╰─ Daenerys Targaryen\n  ├─ The Lannisters of Lannisport\n  ╰─┬─ Tywin Lannister\n    ╰─┬─ Cersei Lannister\n      ├─ Jaime Lannister\n      ╰─ Tyrion Lannister",
		StyleDefault:          "* The Houses of Westeros\n  * The Starks of Winterfell\n    * Eddard Stark\n      * Robb Stark\n      * Sansa Stark\n      * Arya Stark\n      * Bran Stark\n      * Rickon Stark\n    * Lyanna Stark\n    * Benjen Stark\n  * The Targaryens of Dragonstone\n    * Aerys Targaryen\n      * Rhaegar Targaryen\n      * Viserys Targaryen\n      * Daenerys Targaryen\n  * The Lannisters of Lannisport\n    * Tywin Lannister\n      * Cersei Lannister\n      * Jaime Lannister\n      * Tyrion Lannister",
		styleTest:             "^> The Houses of Westeros\nc~f> The Starks of Winterfell\n  T~f> Eddard Stark\n  | T~f> Robb Stark\n  | | i> Sansa Stark\n  | | i> Arya Stark\n  | | i> Bran Stark\n  | | v> Rickon Stark\n  | i> Lyanna Stark\n  | v> Benjen Stark\n  i> The Targaryens of Dragonstone\n  T~f> Aerys Targaryen\n  | c~f> Rhaegar Targaryen\n  |   i> Viserys Targaryen\n  |   v> Daenerys Targaryen\n  i> The Lannisters of Lannisport\n  c~f> Tywin Lannister\n    c~f> Cersei Lannister\n      i> Jaime Lannister\n      v> Tyrion Lannister",
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
