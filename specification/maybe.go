package specification

type Maybe[T any] struct {
	Set   bool
	value T
}

func (m Maybe[T]) Get() (zero T, _ bool) {
	return m.value, m.Set
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		Set:   true,
		value: v,
	}
}

func JustPointer[T any](v *T) Maybe[T] {
	if v == nil {
		return Nothing[T]()
	}
	return Just[T](*v)
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}
