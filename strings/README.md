# `goutils/strings`

[![Go Reference](https://pkg.go.dev/badge/github.com/Shresht7/goutils/strings.svg)](https://pkg.go.dev/github.com/Shresht7/goutils/strings)

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

[![Go Reference](https://pkg.go.dev/badge/github.com/Shresht7/goutils/strings.svg)](https://pkg.go.dev/github.com/Shresht7/goutils/strings)

### `Align`

Aligns the given string to the left, right or center.

```go
func Align(str string, opts *AlignOptions) string
```

Example:

```go
str := strings.Align("ABC", &strings.AlignOptions{
    Width:  9,
    Mode:   strings.AlignCenter,
    Pad: "*",
}) // Output: ***ABC***
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `DetermineIndentation`

Determines the minimum and maximum level of indentation for the given lines.

```go
func DetermineIndentation(lines ...string) (min, max int)
```

Example:

```go
min, max := strings.DetermineIndentation([]string{
    "    Hello, world!",
    "        This is a test.",
    "    Goodbye!",
}...) // Output: 4, 8
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `DetermineWidths`

Determines the width of each column in a 2D array of strings

```go
func DetermineWidths(array [][]string) []int
```

Example:

```go
widths := strings.DetermineWidths([][]string{
    {"Hello", "World!"},
    {"This", "is", "a", "test"},
}) // Output: [5, 6, 1, 4]
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `HereDoc`

HereDoc returns a here-document representation of the given string.

```go
func HereDoc(str string) string
```

Example:

```go
str := strings.HereDoc(`
    Hello, world!
`) // Output: Hello, world!
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `IsSpace`

Checks if the given rune is a space character.

```go
func IsSpace(r rune) bool
```

Example:

```go
fmt.Println(strings.IsSpace(' ')) // true
fmt.Println(strings.IsSpace('\t')) // true
fmt.Println(strings.IsSpace('a')) // false
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

Example:

```go
fmt.Println(strings.PadLeft("ABC", "*", 3)) // ***ABC
fmt.Println(strings.PadRight("ABC", "*", 3)) // ABC***
fmt.Println(strings.Pad("ABC", "*", 2)) // **ABC**
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Dedent`

Removes the indentation from the given string.

```go
func Dedent(s string, n int) string
```

Example:

```go
fmt.Println(strings.Dedent(`
    Hello, world!
        This is a test.
      Goodbye!
`, 2))
// Output:
// Hello, world!
//   This is a test.
// Goodbye!
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📑 License

This project is licensed under the [MIT License](../LICENSE) - see the [LICENSE](../LICENSE) file for details.



<!-- LINKS -->

[top]: #goutils/strings
