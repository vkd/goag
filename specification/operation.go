package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Operation struct {
	PathItem *PathItem

	PathRaw string

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

	Responses Map[Ref[Response]]
}

func NewOperation(pi *PathItem, rawPath string, method httpMethod, operation *openapi3.Operation, specSecurityReqs SecurityRequirements, legacyComponents openapi3.Components, securitySchemes SecuritySchemes, components Components, pathItemParameters openapi3.Parameters, opts SchemaOptions) (*Operation, error) {
	o := &Operation{
		PathItem: pi,

		PathRaw: rawPath,

		Tags:        operation.Tags,
		Summary:     operation.Summary,
		Description: operation.Description,
		OperationID: operation.OperationID,

		HTTPMethod: method.HTTP,
		Method:     method.Title,

		Operation: operation,

		Security: specSecurityReqs,
	}

	var err error
	o.Parameters, err = NewOperationParameters(pathItemParameters, operation.Parameters, components, opts)
	if err != nil {
		return nil, fmt.Errorf("new operation parameters: %w", err)
	}

	if operation.Security != nil {
		o.Security, err = NewSecurityRequirements(*operation.Security, securitySchemes)
		if err != nil {
			return nil, fmt.Errorf("new security requirements: %w", err)
		}
	}

	if operation.RequestBody != nil {
		if operation.RequestBody.Ref != "" {
			v, ok := components.RequestBodies.Get(operation.RequestBody.Ref)
			if !ok {
				return nil, fmt.Errorf("request body %q: not found", operation.RequestBody.Ref)
			}
			o.RequestBody = Just[Ref[RequestBody]](NewRef(v))
		} else {
			o.RequestBody = Just[Ref[RequestBody]](NewRequestBody(operation.RequestBody.Value, components.Schemas, opts))
		}
	}

	o.Responses = NewMap[Ref[Response], *openapi3.ResponseRef](operation.Responses, func(u *openapi3.ResponseRef) Ref[Response] { return nil })
	usedResponses := make(map[*Response]string)
	for i, ro := range o.Responses.List {
		rr := operation.Responses[ro.Name]

		if rr.Ref != "" {
			ref := rr.Ref
			r, ok := components.Responses.Get(ref)
			if !ok {
				return nil, fmt.Errorf("reference %q: not found", ref)
			}
			if usedStatus, ok := usedResponses[r.V.Value()]; ok {
				return nil, fmt.Errorf("the same %q response is used several times (at least for %q and %q responses)", r.Name, usedStatus, ro.Name)
			}
			usedResponses[r.V.Value()] = ro.Name
			r.V.Value().UsedIn = append(r.V.Value().UsedIn, ResponseUsedIn{
				Operation: o,
				Status:    ro.Name,
			})
			o.Responses.List[i].V = NewRef(r)
		} else {
			o.Responses.List[i].V, err = NewResponse(rr.Value, components, opts)
			if err != nil {
				return nil, fmt.Errorf("new response %q: %w", ro.Name, err)
			}
		}
	}

	return o, nil
}

type OperationParameters struct {
	Query   Map[Ref[QueryParameter]]
	Headers Map[Ref[HeaderParameter]]
	Path    Map[Ref[PathParameter]]
	Cookie  Map[Ref[CookieParameter]]
}

func NewOperationParameters(pathParams, operationParams openapi3.Parameters, components Components, opts SchemaOptions) (zero OperationParameters, _ error) {
	out := OperationParameters{
		Query:   NewMapEmpty[Ref[QueryParameter]](0),
		Headers: NewMapEmpty[Ref[HeaderParameter]](0),
		Path:    NewMapEmpty[Ref[PathParameter]](0),
		Cookie:  NewMapEmpty[Ref[CookieParameter]](0),
	}

	for _, param := range append(append(openapi3.Parameters{}, pathParams...), operationParams...) {
		switch param.Value.In {
		case openapi3.ParameterInPath:
			p, err := NewRefPathParam(param, components, opts)
			if err != nil {
				return zero, fmt.Errorf("path param %q: %w", param.Value.Name, err)
			}
			out.Path.Add(p.Value().Name, p)
		case openapi3.ParameterInQuery:
			p, err := NewRefQueryParam(param, components, opts)
			if err != nil {
				return zero, fmt.Errorf("query param %q: %w", param.Value.Name, err)
			}
			out.Query.Add(p.Value().Name, p)
		case openapi3.ParameterInHeader:
			p, err := NewRefHeaderParam(param, components, opts)
			if err != nil {
				return zero, fmt.Errorf("header param %q: %w", param.Value.Name, err)
			}
			out.Headers.Add(p.Value().Name, p)
		case openapi3.ParameterInCookie:
			p, err := NewRefCookieParam(param, components, opts)
			if err != nil {
				return zero, fmt.Errorf("cookie param %q: %w", param.Value.Name, err)
			}
			out.Cookie.Add(p.Value().Name, p)
		}
	}

	return out, nil
}

func NewRefQueryParam(p *openapi3.ParameterRef, components Components, opts SchemaOptions) (Ref[QueryParameter], error) {
	if p.Ref != "" {
		v, ok := components.QueryParameters.Get(p.Ref)
		if !ok {
			return nil, fmt.Errorf("query parameter %q: not found", p.Ref)
		}
		return NewRef(v), nil
	}
	return NewQueryParameter(p.Value, components.Schemas, opts), nil
}

func NewRefHeaderParam(p *openapi3.ParameterRef, components Components, opts SchemaOptions) (Ref[HeaderParameter], error) {
	if p.Ref != "" {
		v, ok := components.HeaderParameters.Get(p.Ref)
		if !ok {
			return nil, fmt.Errorf("header parameter %q: not found", p.Ref)
		}
		return NewRef(v), nil
	}
	return NewHeaderParameter(p.Value, components.Schemas, opts), nil
}

func NewRefPathParam(p *openapi3.ParameterRef, components Components, opts SchemaOptions) (Ref[PathParameter], error) {
	if p.Ref != "" {
		v, ok := components.PathParameters.Get(p.Ref)
		if !ok {
			return nil, fmt.Errorf("path parameter %q: not found", p.Ref)
		}
		return NewRef(v), nil
	}
	return NewPathParameter(p.Value, components.Schemas, opts), nil
}

func NewRefCookieParam(p *openapi3.ParameterRef, components Components, opts SchemaOptions) (Ref[CookieParameter], error) {
	if p.Ref != "" {
		v, ok := components.CookieParameters.Get(p.Ref)
		if !ok {
			return nil, fmt.Errorf("cookie parameter %q: not found", p.Ref)
		}
		return NewRef(v), nil
	}
	return NewCookieParameter(p.Value, components.Schemas, opts), nil
}
