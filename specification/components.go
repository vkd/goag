package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Components struct {
	Schemas ComponentsSchemas

	Headers       Map[Ref[Header]]
	RequestBodies Map[Ref[RequestBody]]
	Responses     Map[Ref[Response]]

	// ---------------- Parameters ----------------
	QueryParameters  Map[Ref[QueryParameter]]
	HeaderParameters Map[Ref[HeaderParameter]]
	PathParameters   Map[Ref[PathParameter]]
	CookieParameters Map[Ref[CookieParameter]]

	SecuritySchemes SecuritySchemes
	Links           Map[Ref[Link]]
	PathItems       Map[Ref[PathItem]]
}

type ComponentsSchemas = Map[Ref[Schema]]
type SecuritySchemes = Map[Ref[SecurityScheme]]

func NewComponents(spec openapi3.Components, opts SchemaOptions) (zero Components, _ error) {
	var cs Components
	var err error

	cs.Schemas, err = NewMapRefSelfSource(spec.Schemas, func(sr *openapi3.SchemaRef, components Sourcer[Schema]) (_ string, zero Ref[Schema], _ error) {
		if sr.Ref != "" {
			return sr.Ref, nil, nil
		}
		return "", NewSchema(sr.Value, components, opts), nil
	}, nil, "#/components/schemas/")
	if err != nil {
		return zero, fmt.Errorf("new schemas: %w", err)
	}

	cs.Headers, err = NewMapRefSelf[Header, *openapi3.HeaderRef](spec.Headers, func(hr *openapi3.HeaderRef) (ref string, _ Ref[Header], _ error) {
		if hr.Ref != "" {
			return hr.Ref, nil, nil
		}
		return "", NewHeader(hr.Value, cs.Schemas, opts), nil
	}, "#/components/headers/")
	if err != nil {
		return zero, fmt.Errorf("new headers: %w", err)
	}

	cs.RequestBodies, err = NewMapRefSelf[RequestBody, *openapi3.RequestBodyRef](spec.RequestBodies, func(hr *openapi3.RequestBodyRef) (ref string, _ Ref[RequestBody], _ error) {
		if hr.Ref != "" {
			return hr.Ref, nil, nil
		}
		return "", NewRequestBody(hr.Value, cs.Schemas, opts), nil
	}, "#/components/requestBodies/")
	if err != nil {
		return zero, fmt.Errorf("new request bodies: %w", err)
	}

	cs.Responses, err = NewMapRefSelf[Response, *openapi3.ResponseRef](spec.Responses, func(rr *openapi3.ResponseRef) (ref string, _ Ref[Response], _ error) {
		if rr.Ref != "" {
			return rr.Ref, nil, nil
		}
		r, err := NewResponse(rr.Value, cs, opts)
		if err != nil {
			return "", r, fmt.Errorf("new response  %w", err)
		}
		return "", r, nil
	}, "#/components/responses/")
	if err != nil {
		return zero, fmt.Errorf("new responses: %w", err)
	}

	queryParameters := make(openapi3.ParametersMap)
	headerParameters := make(openapi3.ParametersMap)
	pathParameters := make(openapi3.ParametersMap)
	cookieParameters := make(openapi3.ParametersMap)
	for k, v := range spec.Parameters {
		switch v.Value.In {
		case "query":
			queryParameters[k] = v
		case "header":
			headerParameters[k] = v
		case "path":
			pathParameters[k] = v
		case "cookie":
			cookieParameters[k] = v
		default:
			return zero, fmt.Errorf("unexpected parameter 'in' value: %q", v.Value.In)
		}
	}

	// ---------------- Parameters ----------------
	cs.QueryParameters, err = NewMapRefSelf[QueryParameter, *openapi3.ParameterRef](queryParameters, func(pr *openapi3.ParameterRef) (ref string, _ Ref[QueryParameter], _ error) {
		if pr.Ref != "" {
			return pr.Ref, nil, nil
		}
		return "", NewQueryParameter(pr.Value, cs.Schemas, opts), nil
	}, "#/components/parameters/")
	if err != nil {
		return zero, fmt.Errorf("new query parameters: %w", err)
	}

	cs.HeaderParameters, err = NewMapRefSelf[HeaderParameter, *openapi3.ParameterRef](headerParameters, func(pr *openapi3.ParameterRef) (ref string, _ Ref[HeaderParameter], _ error) {
		if pr.Ref != "" {
			return pr.Ref, nil, nil
		}
		return "", NewHeaderParameter(pr.Value, cs.Schemas, opts), nil
	}, "#/components/parameters/")
	if err != nil {
		return zero, fmt.Errorf("new header parameters: %w", err)
	}

	cs.PathParameters, err = NewMapRefSelf[PathParameter, *openapi3.ParameterRef](pathParameters, func(pr *openapi3.ParameterRef) (ref string, _ Ref[PathParameter], _ error) {
		if pr.Ref != "" {
			return pr.Ref, nil, nil
		}
		return "", NewPathParameter(pr.Value, cs.Schemas, opts), nil
	}, "#/components/parameters/")
	if err != nil {
		return zero, fmt.Errorf("new path parameters: %w", err)
	}

	cs.CookieParameters, err = NewMapRefSelf[CookieParameter, *openapi3.ParameterRef](cookieParameters, func(pr *openapi3.ParameterRef) (ref string, _ Ref[CookieParameter], _ error) {
		if pr.Ref != "" {
			return pr.Ref, nil, nil
		}
		return "", NewCookieParameter(pr.Value, cs.Schemas, opts), nil
	}, "#/components/parameters/")
	if err != nil {
		return zero, fmt.Errorf("new cookie parameters: %w", err)
	}

	cs.SecuritySchemes, err = NewMapRefSelf[SecurityScheme, *openapi3.SecuritySchemeRef](spec.SecuritySchemes, func(ss *openapi3.SecuritySchemeRef) (ref string, _ Ref[SecurityScheme], _ error) {
		if ss.Ref != "" {
			return ss.Ref, nil, nil
		}
		return "", NewSecurityScheme(ss.Value), nil
	}, "#/components/securitySchemes/")
	if err != nil {
		return zero, fmt.Errorf("new security schemes: %w", err)
	}

	cs.Links, err = NewMapRefSelf[Link, *openapi3.LinkRef](spec.Links, func(lr *openapi3.LinkRef) (ref string, _ Ref[Link], _ error) {
		if lr.Ref != "" {
			return lr.Ref, nil, nil
		}
		return "", NewLink(lr.Value), nil
	}, "#/components/links/")
	if err != nil {
		return zero, fmt.Errorf("new links: %w", err)
	}

	return cs, nil
}
