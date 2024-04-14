package generator

type Maybe[T any] struct {
	Set   bool
	Value T
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		Value: v,
		Set:   true,
	}
}
