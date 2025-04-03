# nullable

## about this package

this package provides simple interface to deal with nullables.
when something is nullable it is unequivocal.

## implementation

```go
package null

type Nullable[T any] struct {
	val T
	ok  bool
}

func New[T any](value T) Nullable[T] {
	return Nullable[T]{
		val: value,
		ok:  true,
	}
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{
		ok: false,
	}
}

func (n Nullable[T]) Ok() (val T, ok bool) {
	return n.val, n.ok
}
```

this is main interface. Rest are just sql drivers and json marshlers.

