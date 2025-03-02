package null

type Nullable[T any] struct {
	Val T
	Ok  bool
}

func New[T any](value T) Nullable[T] {
	return Nullable[T]{
		Val: value,
		Ok:  true,
	}
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{
		Ok: false,
	}
}
