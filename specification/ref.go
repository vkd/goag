package specification

import "fmt"

// type Ref[T any] struct {
// 	Ref *Named[T]
// 	V   *T
// }

type Ref[T any] interface {
	Ref() *Object[string, Ref[T]]
	Value() *T
}

type RefObject[T any] struct {
	V *Object[string, Ref[T]]
}

func NewRefObject[T any](v *Object[string, Ref[T]]) *RefObject[T] {
	return &RefObject[T]{
		V: v,
	}
}
func NewRefObjectSource[T any](key string, source interface {
	Get(string) (*Object[string, Ref[T]], bool)
}) *RefObject[T] {
	v, ok := source.Get(key)
	if !ok {
		panic(fmt.Sprintf("%q key not found: %+v", key, source))
	}
	return &RefObject[T]{
		V: v,
	}
}

func (r *RefObject[T]) Ref() *Object[string, Ref[T]] { return r.V }
func (r *RefObject[T]) Value() *T                    { return r.V.V.Value() }

type NoRef[T any] struct{}

func (NoRef[T]) Ref() *Object[string, Ref[T]] { return nil }

// func NewRef[T any](r R, fn func(*U) T) Ref[T] {
// 	if r.Ref != "" {
// 		var v T
// 		return Ref[T]{Ptr: &v}
// 	}
// 	return Ref[T]{
// 		V: fn(r.Value),
// 	}
// }

// type Ref[T any] interface {
// 	Ref() (name string, ok bool)
// 	Value() *T
// }
