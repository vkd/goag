package generator

import (
	"fmt"
	"strings"

	"github.com/vkd/goag/specification"
)

type Path struct {
	Raw        string
	Dirs       []*PathV
	Parameters specification.Map[*PathV]
}

type PathV struct {
	V          string
	IsVariable bool
	Param      *PathParameter
}

func NewPath(raw string) (zero Path, _ error) {
	p := Path{
		Raw:        raw,
		Parameters: specification.NewMapEmpty[*PathV](0),
	}
	ss := strings.Split(raw, "/")
	if ss[0] != "" {
		return zero, fmt.Errorf("wrong path %q", raw)
	}
	ss = ss[1:]
	for _, dir := range ss {
		d := PathV{V: dir}
		if strings.HasPrefix(dir, "{") && strings.HasSuffix(dir, "}") {
			d = PathV{V: dir[1 : len(dir)-1], IsVariable: true}
		}

		if d.IsVariable {
			p.Parameters.Add(d.V, &d)
		}
		p.Dirs = append(p.Dirs, &d)
	}

	return p, nil
}
