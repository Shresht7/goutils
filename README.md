# `sliceutils`

Higher-order ~~array~~ slice methods in Go! and more.

Contains a set of utility functions to help deal with slices.

> _Inspired by JavaScript's array methods_


---

## 📖 Usage

```go
import (
    "fmt"
    "github.com/Shresht7/sliceutils/slice"
)

func main() {
    s := []int{0, 1, 2, 3}

    slice.ForEach(s, func (value, index int) {
        fmt.Println(index, value)
    })
}
```

To use the method syntax, create a new `Slice`

```go
s := slice.New([]int{0, 1, 2, 3})

s.ForEach(func (value, index int) {
    fmt.Println(index, value)
})
```

Alternatives, just typecast the `[]int` into `Slice[int]` provided by the package.

```go
s := slice.Slice[int]([]int{0, 1, 2, 3})

s.ForEach(func (value, index int) {
    fmt.Println(index, value)
})
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📦 Packages

```sh
go get github.com/Shresht7/sliceutils
```

### `slice`

Contains the utility functions to deal with slices.

```sh
go get github.com/Shresht7/sliceutils/slice
```

### `set`

Contains the utility functions to deal with sets.

```sh
go get github.com/Shresht7/sliceutils/set
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📘 API Reference


### `Equal`

Check if two slices are equal

```go
func Equal(T comparable)(a, b []T) bool
```
```go
a := []int{0, 1, 2, 3}
b := []int{0, 1, 2}
c := []int{0, 1, 2, 3}
slice.Equal(a, b)     //  false
slice.Equal(a, c)     //  true
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `First`, `Last` and `At`

Get the first, last and nth element from the slice

```go
func First[T any](slice []T) T

func Last[T any](slice []T) T

func At[T any](slice []T, pos int) T
```

```go
s := []int{0, 1, 2, 3, 4}
firstElement := slice.First(s)    //  0
lastElement := slice.Last(s)      //  4
thirdElement := slice.At(s, 3)    //  2
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `ForEach`

Iterate over each element of the slice and execute the callback function.

```go
func ForEach[T any](slice []T, cb func(value T, index int))
```

Example:
```go
s := []string{"A", "B", "C", "D", "E"}

slice.ForEach(s, func(value string, index int) {
    fmt.Println("%s is at position %d", value, index)
})
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Filter`

Returns a new slice containing the elements that satisfy the callback.

```go
func Filter[T any](slice []T, cb ReturnCallback[T, bool]) []T
```

Example:
```go
s := []int{0, 1, 2, 3, 4, 5, 6}

filteredElements := slice.Filter(s, func (value, index int) bool {
    return value%2 != 0
})
//  []string{1, 3, 5}
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Map`

Returns a new slice based on the callback criteria.

```go
func Map[From, To any](slice []From, cb ReturnCallback[From, To]) []To
```

Example:
```go
s := []int{0, 1, 2}
newS := slice.Map(s, func (value, index int) int {
    return value * 2
})    //  []int{0, 2, 4}
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Reduce`

Reduces the entire slice into a single value.

```go
func Reduce[From, To any](slice []From, cb ReducerCallback[From, To], initializer To) To
```

Example:
```go
s := []int{0, 1, 2, 3, 4}
sum := slice.Reduce(s, func (accumulator, current, index int, slice []int) int {
    return accumulator + current
}, 0)   //  10
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Every` and `Some`

`Every` returns `true` if the callback is valid for every entry of the slice.

`Some` returns `true` if the callback is valid for any one entry of the slice.

```go
func Every[T any](slice []T, cb ReturnCallback[T, bool]) bool

func Some[T any](slice []T, cb ReturnCallback[T, bool]) bool
```

Example:
```go
s := []int{0, 2, 4, 6, 8}
divisibleBy2 := slice.Every(s, func (value, index int) bool {
    return value % 2 == 0
})    //  true
hasSomethingGreaterThan7 := slice.Some(s, func (value, index int) bool {
    return value > 7
})  //  true
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Find`

Finds the element for which the given callback is `true` and returns a status, the value and the index.

```go
func Find[T any](slice []T, cb ReturnCallback[T, bool]) (bool, T, int)
```

Example:
```go
s := []int{0, 1, 2, 3, 4, 5, 6, 7}
ok, value, index := slice.Find(s, func (v, i int) bool {
    return v == 4
}) //  true, 4, 5
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Includes`

Returns `true` if element is in the slice.

```go
func Includes[T comparable](slice []T, element T) bool
```

Example:
```go
s := []int{0, 1, 2, 3}
a := slice.Includes(s, 3)    //  true
b := slice.Includes(s, 4)    //  false
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Concat`

Concatenates multiple slices into a single slice

```go
func Concat[T any](slices ...[]T) []T
```

Example:
```go
a := []int{0, 1, 2}
b := []int{3, 4}
c := slice.Concat(a, b)   //  []int{0, 1, 2, 3, 4}
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Reverse`

Reverses a slice

```go
func Reverse[T any](slice []T) []T
```

Example:
```go
s := []int{0, 1, 2, 3}
rev := slice.Reverse(s)   //  []int{3, 2, 1, 0}
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Join`

Joins the element of a slice into a string using the given separator.

```go
func Join[T any](slice []T, separator string) string
```

Example:
```go
s := []int{0, 1, 2}
str := slice.Join(s, "-->")   //  "0-->1-->2"
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Push` and `Pop`

`Push` adds elements to the end of the slice.

`Pop` removes elements from the end of the slice.

```go
func Push[T any](slice []T, elements ...T) []T

func Pop[T any](slice *[]T) T
```

Example:
```go
s := []int{0, 1, 2}
s = slice.Push(s, 3)  //  []int{0, 1, 2, 3}
s = slice.Pop(&s)     //  []int{0, 1, 2}
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Shift` and `Unshift`

`Shift` removes elements from the start of the slice.

`Unshift` adds elements to the start of the slice.

```go
func Shift[T any](slice *[]T) T

func Unshift[T any](slice []T, elements ...T) []T
```

Example:
```go
s := []int{0, 1, 2, 3}
s = slice.Shift(&s)       //  []int{0, 1, 2}
s = slice.Unshift(s, 3)   //  []int{0, 1, 2, 3}
```

<div align="right">

⬆️ [Back to Top][top]

</div>

### `Chunk`

Chunks a slice into smaller slices of (at most) given size.

```go
func Chunk[T any](slice []T, size int) [][]T
```

Example:
```go
s := []int{0, 1, 2, 3, 4, 5}
c := slice.Chunk(s, 2)    //  [][]int{{0, 1}, {2, 3}, {4, 5}}
s = slice.Push(s, 6)
d := slice.Chunk(s, 2)    //  [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}
```

<div align="right">

⬆️ [Back to Top][top]

</div>

---

## 📑 License

This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details.





<!-- LINKS -->

[top]: #slice
