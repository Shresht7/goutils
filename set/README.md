# `goutils/set`

[![Go Reference](https://pkg.go.dev/badge/github.com/Shresht7/goutils/set.svg)](https://pkg.go.dev/github.com/Shresht7/goutils/set)

Contains the utility functions to deal with sets.

> _Inspired by JavaScript's set methods_

---

## 📖 Usage

```go
import (
    "fmt"
    "github.com/Shresht7/goutils/set"
)

func main() {
    s := []int{0, 1, 2, 3}

    set.Add(s, 4)
    set.Add(s, 4) // No effect as 4 is already present
    set.Add(s, 5)

    set.ForEach(s, func (value, idx int) {
        fmt.Println(value)
    }) // Output: 0 1 2 3 4 5
}
```

To use the method syntax, create a new `Set`

```go
s := set.New([]int{0, 1, 2, 3})

s.Add(4)
s.Add(4) // No effect as 4 is already present
s.Add(5)

s.ForEach(func (value, idx int) {
    fmt.Println(value)
}) // Output: 0 1 2 3 4 5
```

Alternatives, just typecast the `[]int` into `Set[int]` provided by the package.

```go
s := set.Set[int]([]int{0, 1, 2, 3})

s.Add(4)
s.Add(4) // No effect as 4 is already present
s.Add(5)

s.ForEach(func (value, index int) {
    fmt.Println(index, value)
}) // Output: 0 1 2 3 4 5
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📘 API Reference

[![Go Reference](https://pkg.go.dev/badge/github.com/Shresht7/goutils/set.svg)](https://pkg.go.dev/github.com/Shresht7/goutils/set)

### `Len`

Returns the length of the set.

```go
func (s *Set[T]) Len() int
```

Example:

```go
s := set.New([]int{0, 1, 2, 3})
s.Len() // Output: 4
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Cap`

Returns the capacity of the set.

```go
func (s *Set[T]) Cap() int
```

Example:

```go
s := set.New([]int{0, 1, 2, 3})
s.Cap() // Output: 4
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `AsSlice`

Returns the set as a slice.

```go
func (s *Set[T]) AsSlice() []T
```

Example:

```go
s := set.New([]int{0, 1, 2, 3})
s.AsSlice() // Output: [0 1 2 3]
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Size`

Returns the size of the set.

```go
func (s *Set[T]) Size() int
```

Example:

```go
s := set.New([]int{0, 1, 2, 3})
s.Size() // Output: 4
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Has`

Checks if the set has the given element.

```go
func (s *Set[T]) Has(value T) bool
```

Example:

```go
s := set.New([]int{0, 1, 2, 3})
s.Has(2) // Output: true
s.Has(4) // Output: false
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Add`

Adds an element to the set.

```go
func (s *Set[T]) Add(value T) *Set[T]
```

Example:

```go
s := set.New([]int{0, 1, 2, 3})
s.Add(4)
s.Add(4) // No effect as 4 is already present
s.Add(5)
s.AsSlice() // Output: [0 1 2 3 4 5]
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Delete`

Deletes an element from the set.

```go
func (s *Set[T]) Delete(value T) *Set[T]
```

Example:

```go
s := set.New([]int{0, 1, 2, 3})
s.Delete(2)
s.AsSlice() // Output: [0 1 3]
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Clear`

Clears the set.

```go
func (s *Set[T]) Clear() *Set[T]
```

Example:

```go
s := set.New([]int{0, 1, 2, 3})
s.Clear()
s.AsSlice() // Output: []
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `ForEach`

Iterates over the set and calls the given function for each element.

```go
func (s *Set[T]) ForEach(fn slice.Callback[T])
```

Example:

```go
s := set.New([]int{0, 1, 2, 3})
s.ForEach(func (value, index int) {
    fmt.Println(index, value)
}) // Output: 0 1 2 3
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📑 License

This project is licensed under the [MIT License](../LICENSE) - see the [LICENSE](../LICENSE) file for details.



<!-- LINKS -->

[top]: #goutils/set
