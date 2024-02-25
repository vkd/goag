package specification

import (
	"fmt"
	"sort"
)

type Map[T any] struct {
	List    []*Object[T]
	indexes map[RefKey]*Object[T]
}

type Object[T any] struct {
	Name string
	V    T
}

func NewMapEmpty[T any](size int) Map[T] {
	return Map[T]{
		List:    make([]*Object[T], 0, size),
		indexes: make(map[RefKey]*Object[T], size),
	}
}

func NewMap[T any, U any](m map[string]U, fn func(U) T) Map[T] {
	return NewMapPrefix[T, U](m, fn, "")
}

func NewMapPrefix[T any, U any](m map[string]U, fn func(U) T, prefix RefKey) Map[T] {
	out := NewMapEmpty[T](len(m))

	for k, v := range m {
		out.List = append(out.List, &Object[T]{Name: k, V: fn(v)})
	}
	for i, v := range out.List {
		refKey := prefix + RefKey(v.Name)
		out.indexes[refKey] = out.List[i]
	}
	sort.Slice(out.List, func(i, j int) bool { return out.List[i].Name < out.List[j].Name })
	return out
}

func NewMapRefSelfSource[T any, U any](m map[string]U, fn func(U, Map[Ref[T]]) (ref string, _ Ref[T]), source interface {
	Get(string) (*Object[Ref[T]], bool)
}, prefix RefKey) Map[Ref[T]] {
	out := NewMapPrefix[Ref[T], U](m, func(u U) Ref[T] { return nil }, prefix)
	if source == nil {
		source = out
	}

	for i, o := range out.List {
		ref, v := fn(m[o.Name], out)
		if ref != "" {
			r, ok := source.Get(ref)
			if !ok {
				panic(fmt.Sprintf("reference %q: not found", ref))
			}
			out.List[i].V = NewRefObject(r)
		} else {
			out.List[i].V = v
		}
	}
	return out
}

func NewMapRefSelf[T any, U any](m map[string]U, fn func(U) (ref string, _ Ref[T]), prefix RefKey) Map[Ref[T]] {
	return NewMapRefSelfSource[T, U](m, func(u U, _ Map[Ref[T]]) (ref string, _ Ref[T]) {
		return fn(u)
	}, nil, prefix)
}

func NewMapRefSource[T any, U any](m map[string]U, fn func(U) (ref string, _ Ref[T]), source interface {
	Get(string) (*Object[Ref[T]], bool)
}, prefix RefKey) Map[Ref[T]] {
	return NewMapRefSelfSource[T, U](m, func(u U, m Map[Ref[T]]) (ref string, _ Ref[T]) { return fn(u) }, source, prefix)
}

func (m Map[T]) Get(k string) (*Object[T], bool) {
	if m.indexes == nil {
		return nil, false
	}
	v, ok := m.indexes[RefKey(k)]
	if !ok {
		return nil, false
	}
	return v, true
}

func (m *Map[T]) Add(name string, v T) {
	if m.indexes == nil {
		m.indexes = make(map[RefKey]*Object[T])
	}
	obj := &Object[T]{
		Name: name,
		V:    v,
	}
	m.List = append(m.List, obj)
	m.indexes[RefKey(name)] = obj
}

func (m *Map[T]) Has(k string) bool {
	if m.indexes == nil {
		return false
	}
	_, ok := m.indexes[RefKey(k)]
	return ok
}

type RefKey string
