package specification

import (
	"fmt"
	"strings"
)

type Path struct {
	Spec   string
	Dirs   []*PathDir
	Params PathParameters
}

type PathDir struct {
	Raw      Prefix
	ParamRef *PathParameter
}

func NewPath(s string) (zero Path, _ error) {
	if !strings.HasPrefix(s, "/") {
		return zero, fmt.Errorf("the field name MUST begin with a forward slash (/)")
	}
	out := Path{Spec: s}
	for _, p := range strings.Split(s[1:], "/") {
		prefix := Prefix("/" + p)
		param := &PathParameter{Name: prefix.Name()}
		out.Dirs = append(out.Dirs, &PathDir{
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
