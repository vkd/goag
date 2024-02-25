package generator

import (
	"strings"

	"github.com/vkd/goag/specification"
)

type OperationOld struct {
	Operation *specification.Operation

	Name            OperationName
	HandlerTypeName string

	PathParameters   []Parameter[*specification.PathParameter]
	QueryParameters  []Parameter[*specification.QueryParameter]
	HeaderParameters []Parameter[*specification.HeaderParameter]

	Handler *HandlerOld // Deprecated
}

func NewOperationOld(operation *specification.Operation) *OperationOld {
	o := &OperationOld{
		Operation: operation,
	}
	o.Name = OperationNameOld(operation.OperationID, operation.PathItem.Path, operation.Method)
	o.HandlerTypeName = string(o.Name) + "HandlerFunc"

	for _, pathParam := range operation.Parameters.Path.List {
		o.PathParameters = append(o.PathParameters, Parameter[*specification.PathParameter]{
			Spec: pathParam.V.Value(),

			FieldName: PublicFieldName(pathParam.V.Value().Name),
		})
	}
	for _, query := range operation.Parameters.Query.List {
		o.QueryParameters = append(o.QueryParameters, Parameter[*specification.QueryParameter]{
			Spec: query.V.Value(),

			FieldName: PublicFieldName(query.V.Value().Name),
		})
	}
	for _, header := range operation.Parameters.Headers.List {
		o.HeaderParameters = append(o.HeaderParameters, Parameter[*specification.HeaderParameter]{
			Spec: header.V.Value(),

			FieldName: PublicFieldName(header.V.Value().Name),
		})
	}

	return o
}

func OperationNameOld(operationID string, path specification.PathOld2, method specification.HTTPMethodTitle) OperationName {
	if operationID != "" {
		return OperationName(PublicFieldName(operationID))
	}

	var out string
	for _, dir := range path.Dirs {
		out += PrefixTitle(dir.Raw)
	}

	var suffix string
	if len(path.Dirs) > 1 && path.Dirs[len(path.Dirs)-1].Raw == "/" {
		suffix = "RT"
	}

	return OperationName(string(method) + out + suffix)
}

type Parameter[T interface {
	*specification.PathParameter | *specification.QueryParameter | *specification.HeaderParameter
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
