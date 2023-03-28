# `goutils/set`

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

    set.ForEach(s, func (value, index int) {
        fmt.Println(index, value)
    })
}
```

To use the method syntax, create a new `Set`

```go
s := set.New([]int{0, 1, 2, 3})

s.Add(4)
s.Add(4) // No effect as 4 is already present
s.Add(5)

s.ForEach(func (value, index int) {
    fmt.Println(index, value)
})
```

Alternatives, just typecast the `[]int` into `Set[int]` provided by the package.

```go
s := set.Set[int]([]int{0, 1, 2, 3})

s.Add(4)
s.Add(4) // No effect as 4 is already present
s.Add(5)

s.ForEach(func (value, index int) {
    fmt.Println(index, value)
})
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📘 API Reference

### `Len`

Returns the length of the set.

```go
func (s *Set[T]) Len() int
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Cap`

Returns the capacity of the set.

```go
func (s *Set[T]) Cap() int
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `AsSlice`

Returns the set as a slice.

```go
func (s *Set[T]) AsSlice() []T
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Size`

Returns the size of the set.

```go
func (s *Set[T]) Size() int
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Has`

Checks if the set has the given element.

```go
func (s *Set[T]) Has(value T) bool
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Add`

Adds an element to the set.

```go
func (s *Set[T]) Add(value T) *Set[T]
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Delete`

Deletes an element from the set.

```go
func (s *Set[T]) Delete(value T) *Set[T]
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Clear`

Clears the set.

```go
func (s *Set[T]) Clear() *Set[T]
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `ForEach`

Iterates over the set and calls the given function for each element.

```go
func (s *Set[T]) ForEach(fn slice.Callback[T])
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📑 License

This project is licensed under the [MIT License](../LICENSE) - see the [LICENSE](../LICENSE) file for details.



<!-- LINKS -->

[top]: #goutils/set
