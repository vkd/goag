package generator

type Optional[T any] struct {
	IsSet bool
	Value T
}

func NewOptional[T any](v T) Optional[T] {
	return Optional[T]{
		Value: v,
		IsSet: true,
	}
}
