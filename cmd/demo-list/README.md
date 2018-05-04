Output of `go run cmd/demo-list/demo.go`:

```
A Simple List:
- Game Of Thrones

A Multi-level List:
- Game Of Thrones
--- Winter
  - Is
  - Coming
  --- This
    - Is
    - Known

A List using the Style 'StyleBulletCircle':
● Game Of Thrones
  ● Winter
  ● Is
  ● Coming
    ● This
    ● Is
    ● Known

A List using the Style 'StyleConnectedRounded':
╭─ Game Of Thrones
╰─┬─ Winter
  ├─ Is
  ├─ Coming
  ╰─┬─ This
    ├─ Is
    ╰─ Known

A List using the Style 'funkyStyle':
^> GAME OF THRONES
c<f> WINTER
  i> IS
  i> COMING
  c<f> THIS
    i> IS
    v> KNOWN
```
