package generator

import (
	"strings"

	"github.com/vkd/goag/specification"
)

type Operation struct {
	Operation *specification.Operation

	Name            string
	HandlerTypeName string

	PathParameters   []Parameter[*specification.PathParameter]
	QueryParameters  []Parameter[specification.QueryParameter]
	HeaderParameters []Parameter[*specification.HeaderParameter]

	Handler *HandlerOld // Deprecated
}

func NewOperation(operation *specification.Operation) *Operation {
	o := &Operation{
		Operation: operation,
	}
	o.Name = OperationName(operation.OperationID, operation.PathItem.Path, operation.Method)
	o.HandlerTypeName = o.Name + "HandlerFunc"

	for _, pathParam := range operation.Parameters.Path {
		o.PathParameters = append(o.PathParameters, Parameter[*specification.PathParameter]{
			Spec: pathParam,

			FieldName: PublicFieldName(pathParam.Name),
		})
	}
	for _, query := range operation.Parameters.Query {
		o.QueryParameters = append(o.QueryParameters, Parameter[specification.QueryParameter]{
			Spec: query,

			FieldName: PublicFieldName(query.Name),
		})
	}
	for _, header := range operation.Parameters.Headers {
		o.HeaderParameters = append(o.HeaderParameters, Parameter[*specification.HeaderParameter]{
			Spec: header,

			FieldName: PublicFieldName(header.Name),
		})
	}

	return o
}

func OperationName(operationID string, path specification.Path, method string) string {
	if operationID != "" {
		return PublicFieldName(operationID)
	}

	var out string
	for _, dir := range path.Dirs {
		out += PrefixTitle(dir.Raw)
	}

	var suffix string
	if len(path.Dirs) > 1 && path.Dirs[len(path.Dirs)-1].Raw == "/" {
		suffix = "RT"
	}

	return method + out + suffix
}

type Parameter[T interface {
	*specification.PathParameter | specification.QueryParameter | *specification.HeaderParameter
}] struct {
	Spec T

	FieldName string
	// Type      GoType
}

func PrivateFieldName(name string) string {
	if len(name) == 0 {
		return ""
	}
	return strings.ToLower(name[:1]) + name[1:]
}
