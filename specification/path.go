package specification

import (
	"fmt"
	"strings"
)

// Path always starts with '/'
type Path string

func NewPath(s string) (Path, error) {
	if !strings.HasPrefix(s, "/") {
		return "", fmt.Errorf("path must start with '/'")
	}
	return Path(s), nil
}

func (p Path) Cut() (Prefix, Path, bool) {
	// s := string(p[1:])
	idx := strings.Index(string(p[1:]), "/")
	if idx == -1 {
		return Prefix(p), "", false
	}
	return Prefix(p[:idx+1]), Path(p[idx+1:]), true
}

func (p Path) Name(fn func(Prefix) string, sep string) string {
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
