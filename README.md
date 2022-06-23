# slice-utils

Higher-order ~~array~~ slice methods in Go!

_inspired by JavaScript's array methods_

## Installation

```sh
go get github.com/Shresht7/sliceutils
```

## API Reference

### First

```go
func First[T any](slice []T) T
```

```go
slice := []int{0, 1, 2, 3, 4}
firstElement := First(slice)    //  0
```

### Last

```go
func Last[T any](slice []T) T
```

```go
slice := []int{0, 1, 2, 3, 4}
lastElement := Last(slice)  //  4
```

### Nth

```go
func Nth[T any](slice []T) T
```

```go
slice := []int{0, 1, 2, 3, 4}
secondElement := Nth(slice, 2)  //  1
```

### ForEach

```go
func ForEach[T any](slice []T, cb func(value T, index int))
```

```go
slice := []string{"A", "B", "C", "D", "E"}

ForEach(slice, func(value string, index int) {
    fmt.Println("%s is at position %d", value, index)
})
```
