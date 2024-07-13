package specification

type Ref[T any] interface {
	Ref() *Object[string, Ref[T]]
	Value() *T
}

type refObject[T any] struct {
	V *Object[string, Ref[T]]
}

func NewRef[T any](v *Object[string, Ref[T]]) Ref[T] {
	return &refObject[T]{
		V: v,
	}
}

func (r *refObject[T]) Ref() *Object[string, Ref[T]] { return r.V }
func (r *refObject[T]) Value() *T                    { return r.V.V.Value() }

type NoRef[T any] struct{}

func (NoRef[T]) Ref() *Object[string, Ref[T]] { return nil }
