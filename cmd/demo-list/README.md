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
╰─┬─ Winter
  ├─ Is
  ├─ Coming
  ├─┬─ This
  │ ├─ Is
  │ ╰─ Known
  ├─ The Dark Tower
  ╰─── The Gunslinger

A List using the Style 'funkyStyle':
------------------------------------
^> GAME OF THRONES
c~f> WINTER
  i> IS
  i> COMING
  ~f> THIS
   i> IS
   v> KNOWN
  i> THE DARK TOWER
  c~> THE GUNSLINGER

A List in Markdown format:
--------------------------
[Markdown]   * Game Of Thrones
[Markdown]     * Winter
[Markdown]     * Is
[Markdown]     * Coming
[Markdown]       * This
[Markdown]       * Is
[Markdown]       * Known
[Markdown]     * The Dark Tower
[Markdown]       * The Gunslinger
```
