package spec

import (
	"net/http"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/vkd/goag/generator-v0/source"
)

type Spec struct {
	Swagger *openapi3.Swagger

	Handlers []Handler
}

type Handler struct {
	Path      string
	PathItem  *openapi3.PathItem
	Method    string
	Operation *openapi3.Operation
}

func Parse(spec *openapi3.Swagger) (*Spec, error) {
	var s Spec

	for _, path := range sortedPaths(spec.Paths) {
		pathItem := spec.Paths[path]
		for _, method := range methods() {
			operation := pathItem.GetOperation(method)
			if operation == nil {
				continue
			}
			s.Handlers = append(s.Handlers, Handler{
				Path:      path,
				PathItem:  pathItem,
				Method:    method,
				Operation: operation,
			})
		}
	}
	return &s, nil
}

func (s *Spec) GenerateHandlersFile() (source.HandlersFile, error) {
	panic("not implemented")
}

func sortedPaths(paths openapi3.Paths) (out []string) {
	for k := range paths {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func methods() []string {
	return []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPatch,
		http.MethodPut,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodHead,
		http.MethodOptions,
		http.MethodTrace,
	}
}
