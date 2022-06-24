# slice-utils

Higher-order ~~array~~ slice methods in Go! and more.

Contains a set of utility functions to help deal with slices.

_inspired by JavaScript's array methods_

## 📖 Usage

```sh
go get github.com/Shresht7/sliceutils
```

```go
import (
    "fmt"
    "github.com/Shresht7/sliceutils/slice"
)

func main() {
    s := []int{0, 1, 2, 3}
    ForEach(s, func (value, index int) { fmt.Println(index, value) })
}
```

To use the method syntax, typecast the `[]int` into `Slice[int]` provided by the package.

```go
s := slice.Slice[int]([]int{0, 1, 2, 3})
s.ForEach(func (value, index int) { fmt.Println(index, value) })
```

## 📦 Packages

| package | import                                 |
| ------: | :------------------------------------- |
| `slice` | `github.com/Shresht7/sliceutils/slice` |
|   `set` | `github.com/Shresht7/sliceutils/set`   |

---

## API Reference

🚧 Work in Progress 🚧

### `Equal`

Check if two slices are equal

```go
func Equal(T comparable)(a, b []T) bool
```
```go
a := []int{0, 1, 2, 3}
b := []int{0, 1, 2}
c := []int{0, 1, 2, 3}
Equal(a, b)     //  false
Equal(a, c)     //  true
```

### `First`, `Last` and `Nth`

Get the first, last and nth element from the slice

```go
func First[T any](slice []T) T

func Last[T any](slice []T) T

func Nth[T any](slice []T) T
```

```go
slice := []int{0, 1, 2, 3, 4}
firstElement := First(slice)    //  0
lastElement := Last(slice)      //  4
thirdElement := Nth(3)          //  2
```

### `ForEach`

Iterate over each element of the slice and execute the callback function.

```go
func ForEach[T any](slice []T, cb func(value T, index int))
```

Example:
```go
slice := []string{"A", "B", "C", "D", "E"}

ForEach(slice, func(value string, index int) {
    fmt.Println("%s is at position %d", value, index)
})
```

### `Filter`

Returns a new slice containing the elements that satisfy the callback.

```go
func Filter[T any](slice []T, cb ReturnCallback[T, bool]) []T
```

Example:
```go
slice := []int{0, 1, 2, 3, 4, 5, 6}
filteredElements := Filter(slice, func (value, index int) bool { return value%2 != 0 })      //  []string{1, 3, 5}
```

### `Map`

Returns a new slice based on the callback criteria.

```go
func Map[From, To any](slice []From, cb ReturnCallback[From, To]) []To
```

Example:
```go
slice := []int{0, 1, 2}
newSlice := Map(slice, func (value, index int) int { return value * 2 })    //  []int{0, 2, 4}
```

### `Reduce`

Reduces the entire slice into a single value.

```go
func Reduce[From, To any](slice []From, cb ReducerCallback[From, To], initializer To) To
```

Example:
```go
slice := []int{0, 1, 2, 3, 4}
sum := Reduce(slice, func (accumulator, current, index int, slice []int) int { return accumulator + current }, 0)   //  10
```

### `Every` and `Some`

`Every` returns `true` if the callback is valid for every entry of the slice.

`Some` returns `true` if the callback is valid for any one entry of the slice.

```go
func Every[T any](slice []T, cb ReturnCallback[T, bool]) bool

func Some[T any](slice []T, cb ReturnCallback[T, bool]) bool
```

Example:
```go
slice := []int{0, 2, 4, 6, 8}
divisibleBy2 := Every(slice, func (value, index int) bool { return value % 2 == 0 })    //  true
hasSomethingGreaterThan7 := Some(slice, func (value, index int) bool { return value > 7 })  //  true
```

### `Find`

Finds the element for which the given callback is `true` and returns a status, the value and the index.

```go
func Find[T any](slice []T, cb ReturnCallback[T, bool]) (bool, T, int)
```

Example:
```go
slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
ok, value, index := Find(slice, func (v, i int) bool { return v == 4 }) //  true, 4, 5
```

### `Includes`

Returns `true` if element is in the slice.

```go
func Includes[T comparable](slice []T, element T) bool
```

Example:
```go
slice := []int{0, 1, 2, 3}
a := Includes(slice, 3)    //  true
b := Includes(slice, 4)    //  false
```

### `Concat`

Concatenates multiple slices into a single slice

```go
func Concat[T any](slices ...[]T) []T
```

Example:
```go
a := []int{0, 1, 2}
b := []int{3, 4}
c := Concat(a, b)   //  []int{0, 1, 2, 3, 4}
```

### `Reverse`

Reverses a slice

```go
func Reverse[T any](slice []T) []T
```

Example:
```go
slice := []int{0, 1, 2, 3}
rev := Reverse(slice)   //  []int{3, 2, 1, 0}
```

### `Join`

Joins the element of a slice into a string using the given separator.

```go
func Join[T any](slice []T, separator string) string
```

Example:
```go
slice := []int{0, 1, 2}
str := Join(slice, "-->")   //  "0-->1-->2"
```

### `Push` and `Pop`

`Push` adds elements to the end of the slice.

`Pop` removes elements from the end of the slice.

```go
func Push[T any](slice []T, elements ...T) []T

func Pop[T any](slice *[]T) T
```

Example:
```go
slice := []int{0, 1, 2}
slice = Push(slice, 3)  //  []int{0, 1, 2, 3}
slice = Pop(&slice)     //  []int{0, 1, 2}
```

### `Shift` and `Unshift`

`Shift` removes elements from the start of the slice.

`Unshift` adds elements to the start of the slice.

```go
func Shift[T any](slice *[]T) T

func Unshift[T any](slice []T, elements ...T) []T
```

Example:
```go
slice := []int{0, 1, 2, 3}
slice = Shift(&slice)       //  []int{0, 1, 2}
slice = Unshift(slice, 3)   //  []int{0, 1, 2, 3}
```

### Chunk

Chunks a slice into smaller slices of (at most) given size.

```go
func Chunk[T any](slice []T, size int) [][]T
```

Example:
```go
slice := []int{0, 1, 2, 3, 4, 5}
c := Chunk(slice, 2)    //  [][]int{{0, 1}, {2, 3}, {4, 5}}
slice = Push(slice, 6)
c := Chunk(slice, 2)    //  [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}
```

---

## 📑 License

> [MIT License](./LICENSE)
