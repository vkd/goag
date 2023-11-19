package generator

type Optional[T any] struct {
	Value T
	OK    bool
}

func (o *Optional[T]) Set(v T) {
	o.Value = v
	o.OK = true
}
