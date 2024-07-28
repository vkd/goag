package specification

import (
	"fmt"
	"sort"

	"golang.org/x/exp/maps"
)

type Map[T any] struct {
	List    []*Object[string, T]
	indexes map[string]*Object[string, T]
}

type Object[K, T any] struct {
	Name K
	V    T
}

func NewMapEmpty[T any](size int) Map[T] {
	return Map[T]{
		List:    make([]*Object[string, T], 0, size),
		indexes: make(map[string]*Object[string, T], size),
	}
}

func NewMap[T any, U any](m map[string]U, fn func(U) (T, error)) (Map[T], error) {
	return NewMapPrefix[T, U](m, fn, "")
}

func NewMapPrefix[T any, U any](m map[string]U, fn func(U) (T, error), prefix string) (zero Map[T], _ error) {
	out := NewMapEmpty[T](len(m))

	keys := maps.Keys(m)
	sort.Strings(keys)

	for _, k := range keys {
		value, err := fn(m[k])
		if err != nil {
			return zero, fmt.Errorf("%q key: %w", k, err)
		}
		o := &Object[string, T]{Name: k, V: value}
		out.List = append(out.List, o)
		out.indexes[prefix+k] = o
	}
	return out, nil
}

type Sourcer[T any] interface {
	Get(string) (*Object[string, Ref[T]], bool)
}

type SourcerFunc[T any] func(string) (*Object[string, Ref[T]], bool)

func (s SourcerFunc[T]) Get(str string) (*Object[string, Ref[T]], bool) { return s(str) }

func NewMapRefSelfSource[T any, U any](m map[string]U, fn func(U, Sourcer[T]) (ref string, _ Ref[T], _ error), source Sourcer[T], prefix string) (zero Map[Ref[T]], _ error) {
	out, err := NewMapPrefix[Ref[T], U](m, func(u U) (Ref[T], error) { return nil, nil }, prefix)
	if err != nil {
		return zero, fmt.Errorf("new empty map: %w", err)
	}

	if source == nil {
		source = out
	}

	for i, o := range out.List {
		ref, v, err := fn(m[o.Name], out)
		if err != nil {
			return zero, fmt.Errorf("map key %q: %w", o.Name, err)
		}
		if ref != "" {
			r, ok := source.Get(ref)
			if !ok {
				return zero, fmt.Errorf("reference %q: not found", ref)
			}
			out.List[i].V = NewRef(r)
		} else {
			out.List[i].V = v
		}
	}
	return out, nil
}

func NewMapRefSelf[T any, U any](m map[string]U, fn func(U) (ref string, _ Ref[T], _ error), prefix string) (Map[Ref[T]], error) {
	return NewMapRefSelfSource[T, U](m, func(u U, _ Sourcer[T]) (ref string, _ Ref[T], _ error) {
		return fn(u)
	}, nil, prefix)
}

func (m Map[T]) Get(k string) (*Object[string, T], bool) {
	if m.indexes == nil {
		return nil, false
	}
	v, ok := m.indexes[k]
	if !ok {
		return nil, false
	}
	return v, true
}

func (m *Map[T]) Add(name string, v T) {
	if m.indexes == nil {
		m.indexes = make(map[string]*Object[string, T])
	}
	if o, ok := m.indexes[name]; ok {
		o.V = v
		return
	}
	obj := &Object[string, T]{
		Name: name,
		V:    v,
	}
	m.List = append(m.List, obj)
	m.indexes[name] = obj
}

func (m *Map[T]) Has(k string) bool {
	if m.indexes == nil {
		return false
	}
	_, ok := m.indexes[k]
	return ok
}
