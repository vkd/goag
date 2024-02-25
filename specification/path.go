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

// PathOld always starts with '/'
type PathOld string

func NewPathOld(s string) (PathOld, error) {
	if !strings.HasPrefix(s, "/") {
		return "", fmt.Errorf("path must start with '/'")
	}
	return PathOld(s), nil
}

func (p PathOld) Cut() (Prefix, PathOld, bool) {
	// s := string(p[1:])
	idx := strings.Index(string(p[1:]), "/")
	if idx == -1 {
		return Prefix(p), "", false
	}
	return Prefix(p[:idx+1]), PathOld(p[idx+1:]), true
}

func (p PathOld) Name(fn func(Prefix) string, sep string) string {
	var out string
	for {
		prefix, path, ok := p.Cut()
		if !ok {
			return out + fn(prefix)
		}
		out += fn(prefix) + sep
		p = path
	}
}

func (p PathOld) String() string { return string(p) }

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
