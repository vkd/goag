package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Operation struct {
	*specification.Operation

	Name OperationName
	Path specification.PathOld2

	APIHandlerFieldName string
	HandlerTypeName     string

	RequestTypeName  string
	ResponseTypeName string

	Params OperationParams

	DefaultResponse Optional[*Response]
	Responses       []*Response
}

func NewOperation(s *specification.Operation, components specification.Components) (zero *Operation, _ Imports, _ error) {
	name := NewOperationName(s)
	o := Operation{
		Operation: s,
		Name:      name,
		Path:      s.PathItem.Path,

		APIHandlerFieldName: string(name) + "Handler",
		HandlerTypeName:     string(name) + "HandlerFunc",

		RequestTypeName:  string(name) + "Params",
		ResponseTypeName: string(name) + "Response",
	}

	var imports Imports
	var err error
	o.Params, imports, err = NewOperationParams(s.Parameters)
	if err != nil {
		return zero, nil, fmt.Errorf("new operation params: %w", err)
	}

	if s.DefaultResponse != nil {
		o.DefaultResponse = NewOptional[*Response](NewResponse(name, s.DefaultResponse))
	}
	for _, r := range s.Responses {
		o.Responses = append(o.Responses, NewResponse(name, r))
	}

	return &o, imports, nil
}

type OperationName string

func NewOperationName(s *specification.Operation) OperationName {
	if s.OperationID != "" {
		return OperationName(PublicFieldName(s.OperationID))
	}

	path := s.PathItem.Path

	var out string
	for _, dir := range path.Dirs {
		out += PrefixTitle(dir.Raw)
	}

	var suffix string
	if len(path.Dirs) > 1 && path.Dirs[len(path.Dirs)-1].Raw == "/" {
		suffix = "RT"
	}

	return OperationName(string(s.Method) + out + suffix)
}

type OperationParams struct {
	Query   specification.Map[*QueryParameter]
	Headers specification.Map[*HeaderParameter]
	Path    specification.Map[*PathParameter]
	Cookie  specification.Map[*CookieParameter]
}

func NewOperationParams(params specification.OperationParameters) (zero OperationParams, _ Imports, _ error) {
	var op OperationParams
	var imports Imports

	for _, p := range params.Query.List {
		param, ims, err := NewQueryParameter(p.V.Value())
		if err != nil {
			return zero, nil, fmt.Errorf("new query parameter: %w", err)
		}
		op.Query.Add(p.Name, param)
		imports = append(imports, ims...)
	}

	for _, p := range params.Headers.List {
		param, ims, err := NewHeaderParameter(p.V.Value())
		if err != nil {
			return zero, nil, fmt.Errorf("new header parameter: %w", err)
		}
		op.Headers.Add(p.Name, param)
		imports = append(imports, ims...)
	}

	for _, p := range params.Path.List {
		param, ims, err := NewPathParameter(p.V)
		if err != nil {
			return zero, nil, fmt.Errorf("new path parameter: %w", err)
		}
		op.Path.Add(p.Name, param)
		imports = append(imports, ims...)
	}

	return op, imports, nil
}
