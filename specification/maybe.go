package specification

type Maybe[T any] struct {
	Set   bool
	Value T
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		Set:   true,
		Value: v,
	}
}

func NewMaybe[T any](v *T) Maybe[T] {
	if v == nil {
		return Nothing[T]()
	}
	return Just[T](*v)
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}
