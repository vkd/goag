package generator

type Maybe[T any] struct {
	IsSet bool
	Value T
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		Value: v,
		IsSet: true,
	}
}

func Nothing[T any]() Maybe[T] { return Maybe[T]{IsSet: false} }

func (m Maybe[T]) Get() (zero T, _ bool) {
	return m.Value, m.IsSet
}

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}
