package generator

type Optional[T any] struct {
	Value T
	OK    bool
}

func NewOptional[T any](v T) Optional[T] {
	return Optional[T]{
		Value: v,
		OK:    true,
	}
}

func (o *Optional[T]) Set(v T) {
	o.Value = v
	o.OK = true
}
