package generator

import (
	"github.com/vkd/goag/specification"
)

type MappedList[I comparable, O any] struct {
	List []O
	m    map[I]*O
}

func NewMappedList[I comparable, O any](m specification.Map[I]) *MappedList[I, O] {
	out := &MappedList[I, O]{
		List: make([]O, len(m.List)),
		m:    make(map[I]*O, len(m.List)),
	}
	for i, c := range m.List {
		out.m[c.V] = &out.List[i]
	}
	return out
}

func (m MappedList[I, O]) Get(v I) (*O, bool) {
	out, ok := m.m[v]
	return out, ok
}
