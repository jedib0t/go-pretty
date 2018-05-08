Output of `go run cmd/demo-list/demo.go`:

```
A Simple List:
--------------
* Game Of Thrones
* The Dark Tower

A Multi-level List:
-------------------
* Game Of Thrones
  * Winter
  * Is
  * Coming
    * This
    * Is
    * Known
* The Dark Tower
  * The Gunslinger

A List using the Style 'StyleBulletCircle':
-------------------------------------------
● Game Of Thrones
  ● Winter
  ● Is
  ● Coming
    ● This
    ● Is
    ● Known
● The Dark Tower
  ● The Gunslinger

A List using the Style 'StyleConnectedRounded':
-----------------------------------------------
╭─ Game Of Thrones
├─┬─ Winter
│ ├─ Is
│ ├─ Coming
│ ╰─┬─ This
│   ├─ Is
│   ╰─ Known
├─ The Dark Tower
╰─── The Gunslinger

A List using the Style 'funkyStyle':
------------------------------------
^> GAME OF THRONES
~f> WINTER
 i> IS
 i> COMING
 c~f> THIS
   i> IS
   v> KNOWN
i> THE DARK TOWER
c~> THE GUNSLINGER

A List in HTML format:
----------------------
[HTML] <ul class="go-pretty-table">
[HTML]   <li>Game Of Thrones</li>
[HTML]   <ul class="go-pretty-table-1">
[HTML]     <li>Winter</li>
[HTML]     <li>Is</li>
[HTML]     <li>Coming</li>
[HTML]     <ul class="go-pretty-table-2">
[HTML]       <li>This</li>
[HTML]       <li>Is</li>
[HTML]       <li>Known</li>
[HTML]     </ul>
[HTML]   </ul>
[HTML]   <li>The Dark Tower</li>
[HTML]   <ul class="go-pretty-table-1">
[HTML]     <li>The Gunslinger</li>
[HTML]   </ul>
[HTML] </ul>

A List in Markdown format:
--------------------------
[Markdown]   * Game Of Thrones
[Markdown]     * Winter
[Markdown]     * Is
[Markdown]     * Coming
[Markdown]       * This
[Markdown]       * Is
[Markdown]       * Known
[Markdown]   * The Dark Tower
[Markdown]     * The Gunslinger

```
