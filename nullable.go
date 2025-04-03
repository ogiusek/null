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
