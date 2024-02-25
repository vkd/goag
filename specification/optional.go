package specification

type Optional[T any] struct {
	IsSet bool
	Value T
}

func NewOptional[T any](v T) Optional[T] {
	return Optional[T]{
		IsSet: true,
		Value: v,
	}
}

func NewOptionalPtr[T any](v *T) Optional[T] {
	if v == nil {
		return NewOptionalEmpty[T]()
	}
	return NewOptional[T](*v)
}

func NewOptionalEmpty[T any]() Optional[T] {
	return Optional[T]{}
}
