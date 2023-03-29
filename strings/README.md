# `goutils/strings`

A collection of utility functions for working with strings in Go.

---

## 📖 Usage

```go
import (
    "fmt"
    "github.com/Shresht7/goutils/strings"
)

func main() {
    
    fmt.Println(strings.PadLeft("ABC", "*", 3)) // ***ABC***
    
    fmt.Println(strings.HereDoc(`
        Hello, world!
    `)) // Hello, world!
}
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📘 API Reference

### `Align`

Aligns the given string to the left, right or center.

```go
func Align(str string, opts *AlignOptions) string
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `DetermineIndentation`

Determines the minimum and maximum level of indentation for the given lines.

```go
func DetermineIndentation(lines []string) (min, max int)
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `DetermineWidths`

Determines the width of each column in a 2D array of strings

```go
func DetermineWidths(array [][]string) []int
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `HereDoc`

HereDoc returns a here-document representation of the given string.

```go
func HereDoc(str string) string
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `IsSpace`

Checks if the given rune is a space character.

```go
func IsSpace(r rune) bool
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Pad`

Pads the given string with the given character to the left, right or center.

```go
func Pad(s, char string, count int) string
func PadLeft(s, char string, count int) string
func PadRight(s, char string, count int) string
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `RemoveIndentation`

Removes the indentation from the given string.

```go
func RemoveIndentation(n int, lines []string) []string
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Wrap`

Wraps the given string to the given width.

```go
func Wrap(s, str string) string
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📑 License

This project is licensed under the [MIT License](../LICENSE) - see the [LICENSE](../LICENSE) file for details.



<!-- LINKS -->

[top]: #goutils/strings
