package generator

type Maybe[T any] struct {
	Set   bool
	value T
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		value: v,
		Set:   true,
	}
}

func (m Maybe[T]) Get() (zero T, _ bool) {
	return m.value, m.Set
}
