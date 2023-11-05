package specification

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

type Spec struct {
	Swagger *openapi3.Swagger

	Paths      []*PathItem
	Operations []*Operation
}

func ParseSwagger(spec *openapi3.Swagger) (*Spec, error) {
	var s Spec

	for _, path := range sortedPaths(spec.Paths) {
		p, err := NewPath(path)
		if err != nil {
			return nil, fmt.Errorf("validate path %q: %w", path, err)
		}
		pathItem := spec.Paths[path]
		pi := &PathItem{
			Path:     p,
			PathItem: pathItem,
		}
		for _, method := range methods() {
			operation := pathItem.GetOperation(method.HTTP)
			if operation == nil {
				continue
			}
			o := &Operation{
				Path:       p,
				PathItem:   pathItem,
				HTTPMethod: method.HTTP,
				Method:     method.Title,
				Operation:  operation,
			}
			pi.Operations = append(pi.Operations, o)
			s.Operations = append(s.Operations, o)
		}
		s.Paths = append(s.Paths, pi)
	}
	return &s, nil
}

type PathItem struct {
	Path     Path
	PathItem *openapi3.PathItem

	Operations []*Operation
}

type Operation struct {
	Path       Path
	PathItem   *openapi3.PathItem
	HTTPMethod string
	Method     string
	Operation  *openapi3.Operation
}

func sortedPaths(paths openapi3.Paths) (out []string) {
	for k := range paths {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func methods() []method {
	return []method{
		{http.MethodGet, "Get"},
		{http.MethodPost, "Post"},
		{http.MethodPatch, "Patch"},
		{http.MethodPut, "Put"},
		{http.MethodDelete, "Delete"},
		{http.MethodConnect, "Connect"},
		{http.MethodHead, "Head"},
		{http.MethodOptions, "Options"},
		{http.MethodTrace, "Trace"},
	}
}

type method struct {
	HTTP  string
	Title string
}
