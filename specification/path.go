package specification

import (
	"fmt"
	"strings"
)

type Path struct {
	Raw  string
	Dirs []*PathDir
	Refs Map[*PathDir]
}

type PathDir struct {
	V          string
	IsVariable bool
	Param      Ref[PathParameter]
}

func NewPath(raw string) Path {
	p := Path{
		Raw:  raw,
		Refs: NewMapEmpty[*PathDir](0),
	}
	ss := strings.Split(raw, "/")
	if ss[0] != "" {
		panic(fmt.Sprintf("wrong path %q", raw))
	}
	ss = ss[1:]
	for _, dir := range ss {
		d := PathDir{V: dir}
		if strings.HasPrefix(dir, "{") && strings.HasSuffix(dir, "}") {
			d = PathDir{V: dir[1 : len(dir)-1], IsVariable: true}
		}

		if d.IsVariable {
			p.Refs.Add(d.V, &d)
		}
		p.Dirs = append(p.Dirs, &d)
	}

	return p
}

type PathOld2 struct {
	Spec   string
	Dirs   []*PathDirOld
	Params PathParameters
}

type PathDirOld struct {
	Raw      Prefix
	ParamRef *PathParameterOld
}

func NewPathOld2(s string) (zero PathOld2, _ error) {
	if !strings.HasPrefix(s, "/") {
		return zero, fmt.Errorf("the field name MUST begin with a forward slash (/)")
	}
	out := PathOld2{Spec: s}
	for _, p := range strings.Split(s[1:], "/") {
		prefix := Prefix("/" + p)
		param := &PathParameterOld{Name: prefix.Name()}
		out.Dirs = append(out.Dirs, &PathDirOld{
			Raw:      prefix,
			ParamRef: param,
		})
		out.Params = append(out.Params, param)
	}

	return out, nil
}

// Prefix always starts with '/'
type Prefix string

func (p Prefix) IsVariable() bool {
	return strings.HasPrefix(string(p), "/{") && strings.HasSuffix(string(p), "}")
}

func (p Prefix) Name() string {
	if p.IsVariable() {
		return string(p[2 : len(p)-1])
	}
	return string(p[1:])
}
