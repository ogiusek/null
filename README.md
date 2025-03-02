# nullable

this package is basically this:
```go
package null

type Nullable[T any] struct {
	Val T
	Ok  bool
}
```

with these constructors:
```go
func New[T any](value T) Nullable[T] {
	return Nullable[T]{
		Val: value,
		Ok:  true,
	}
}
```
and
```go
func Null[T any]() Nullable[T] {
	return Nullable[T]{
		Ok: false,
	}
}
```

other files are json marshlers and sql drivers.
