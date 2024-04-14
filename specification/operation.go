package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Operation struct {
	PathItem *PathItem

	Path Path

	Tags        []string
	Summary     string
	Description string
	OperationID string

	Parameters OperationParameters

	RequestBody Maybe[Ref[RequestBody]]

	// Deprecated // TODO

	HTTPMethod HTTPMethod
	Method     HTTPMethodTitle

	Operation *openapi3.Operation

	Security SecurityRequirements

	// DefaultResponse Optional[Ref[Response]]
	// Responses       []*ResponseOld

	Responses Map[Ref[Response]]
}

func NewOperation(pi *PathItem, rawPath string, method httpMethod, operation *openapi3.Operation, specSecurityReqs SecurityRequirements, legacyComponents openapi3.Components, securitySchemes SecuritySchemes, components Components) (*Operation, error) {
	o := &Operation{
		PathItem: pi,

		Path: NewPath(rawPath),

		Tags:        operation.Tags,
		Summary:     operation.Summary,
		Description: operation.Description,
		OperationID: operation.OperationID,

		Parameters: NewOperationParameters(pi.PathItem.Parameters, operation.Parameters, components),

		HTTPMethod: method.HTTP,
		Method:     method.Title,

		Operation: operation,

		Security: specSecurityReqs,
	}

	err := o.mapPathParams()
	if err != nil {
		return nil, fmt.Errorf("map path parameters: %w", err)
	}

	if operation.Security != nil {
		o.Security = NewSecurityRequirements(*operation.Security, securitySchemes)
	}

	if operation.RequestBody != nil {
		if operation.RequestBody.Ref != "" {
			o.RequestBody = Just[Ref[RequestBody]](NewRefObjectSource[RequestBody](operation.RequestBody.Ref, components.RequestBodies))
		} else {
			o.RequestBody = Just[Ref[RequestBody]](NewRequestBody(operation.RequestBody.Value, components.Schemas))
		}
	}

	o.Responses = NewMapPrefix[Ref[Response], *openapi3.ResponseRef](operation.Responses, func(u *openapi3.ResponseRef) Ref[Response] { return nil }, "")
	usedResponses := make(map[*Response]string)
	for i, ro := range o.Responses.List {
		rr := operation.Responses[ro.Name]

		if rr.Ref != "" {
			ref := rr.Ref
			r, ok := components.Responses.Get(ref)
			if !ok {
				panic(fmt.Sprintf("reference %q: not found", ref))
			}
			if usedStatus, ok := usedResponses[r.V.Value()]; ok {
				return nil, fmt.Errorf("the same %q response is used several times (at least for %q and %q responses)", r.Name, usedStatus, ro.Name)
			}
			usedResponses[r.V.Value()] = ro.Name
			r.V.Value().UsedIn = append(r.V.Value().UsedIn, ResponseUsedIn{
				Operation: o,
				Status:    ro.Name,
			})
			o.Responses.List[i].V = NewRefObject(r)
		} else {
			o.Responses.List[i].V = NewResponse(rr.Value, components)
		}
	}

	return o, nil
}

func (o *Operation) mapPathParams() error {
	for _, pp := range o.Parameters.Path.List {
		obj, ok := o.Path.Refs.Get(pp.Name)
		if !ok {
			return fmt.Errorf("%q path parameter: not found in %q endpoint", pp.Name, o.Path.Raw)
		}
		obj.V.Param = pp.V
	}
	for _, pp := range o.Path.Refs.List {
		if pp.V.IsVariable && pp.V.Param == nil {
			return fmt.Errorf("%q endpoint: %q param is not defined", o.Path.Raw, pp.V.V)
		}
	}
	return nil
}

type OperationParameters struct {
	Query   Map[Ref[QueryParameter]]
	Headers Map[Ref[HeaderParameter]]
	Path    Map[Ref[PathParameter]]
	Cookie  Map[Ref[CookieParameter]]
}

func NewOperationParameters(pathParams, operationParams openapi3.Parameters, components Components) OperationParameters {
	out := OperationParameters{
		Query:   NewMapEmpty[Ref[QueryParameter]](0),
		Headers: NewMapEmpty[Ref[HeaderParameter]](0),
		Path:    NewMapEmpty[Ref[PathParameter]](0),
		Cookie:  NewMapEmpty[Ref[CookieParameter]](0),
	}

	for _, param := range append(append(openapi3.Parameters{}, pathParams...), operationParams...) {
		switch param.Value.In {
		case openapi3.ParameterInPath:
			p := NewRefPathParam(param, components)
			out.Path.Add(p.Value().Name, p)
		case openapi3.ParameterInQuery:
			p := NewRefQueryParam(param, components)
			out.Query.Add(p.Value().Name, p)
		case openapi3.ParameterInHeader:
			p := NewRefHeaderParam(param, components)
			out.Headers.Add(p.Value().Name, p)
		case openapi3.ParameterInCookie:
			p := NewRefCookieParam(param, components)
			out.Cookie.Add(p.Value().Name, p)
		}
	}

	return out
}

func NewRefQueryParam(p *openapi3.ParameterRef, components Components) Ref[QueryParameter] {
	if p.Ref != "" {
		return NewRefObjectSource[QueryParameter](p.Ref, components.QueryParameters)
	}
	return NewQueryParameter(p.Value, components.Schemas)
}

func NewRefHeaderParam(p *openapi3.ParameterRef, components Components) Ref[HeaderParameter] {
	if p.Ref != "" {
		return NewRefObjectSource[HeaderParameter](p.Ref, components.HeaderParameters)
	}
	return NewHeaderParameter(p.Value, components.Schemas)
}

func NewRefPathParam(p *openapi3.ParameterRef, components Components) Ref[PathParameter] {
	if p.Ref != "" {
		return NewRefObjectSource[PathParameter](p.Ref, components.PathParameters)
	}
	return NewPathParameter(p.Value, components.Schemas)
}

func NewRefCookieParam(p *openapi3.ParameterRef, components Components) Ref[CookieParameter] {
	if p.Ref != "" {
		return NewRefObjectSource[CookieParameter](p.Ref, components.CookieParameters)
	}
	return NewCookieParameter(p.Value, components.Schemas)
}
