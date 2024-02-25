package specification

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type Components struct {
	Schemas   ComponentsSchemas
	Responses Map[Ref[Response]]

	// ---------------- Parameters ----------------
	QueryParameters  Map[Ref[QueryParameter]]
	HeaderParameters Map[Ref[HeaderParameter]]
	PathParameters   Map[Ref[PathParameter]]
	CookieParameters Map[Ref[CookieParameter]]

	RequestBodies   Map[Ref[RequestBody]]
	Headers         Map[Ref[Header]]
	SecuritySchemes SecuritySchemes
	Links           Map[Ref[Link]]
	PathItems       Map[Ref[PathItem]]
}

type ComponentsSchemas = Map[Ref[Schema]]
type SecuritySchemes = Map[Ref[SecurityScheme]]

func NewComponents(spec openapi3.Components) Components {
	var cs Components

	cs.Schemas = NewMapRefSelfSource(spec.Schemas, func(sr *openapi3.SchemaRef, components ComponentsSchemas) (_ string, zero Ref[Schema]) {
		if sr.Ref != "" {
			return sr.Ref, nil
		}
		return "", NewSchema(sr.Value, components)
	}, nil, "#/components/schemas/")

	cs.Responses = NewMapRefSelf[Response, *openapi3.ResponseRef](spec.Responses, func(rr *openapi3.ResponseRef) (ref string, _ Ref[Response]) {
		if rr.Ref != "" {
			return rr.Ref, nil
		}
		return "", NewResponse(rr.Value, cs)
	}, "#/components/responses/")

	// ---------------- Parameters ----------------
	cs.QueryParameters = NewMapRefSelf[QueryParameter, *openapi3.ParameterRef](spec.Parameters, func(pr *openapi3.ParameterRef) (ref string, _ Ref[QueryParameter]) {
		if pr.Ref != "" {
			return pr.Ref, nil
		}
		return "", NewQueryParameter(pr.Value, cs.Schemas)
	}, "#/components/parameters/")

	cs.HeaderParameters = NewMapRefSelf[HeaderParameter, *openapi3.ParameterRef](spec.Parameters, func(pr *openapi3.ParameterRef) (ref string, _ Ref[HeaderParameter]) {
		if pr.Ref != "" {
			return pr.Ref, nil
		}
		return "", NewHeaderParameter(pr.Value, cs.Schemas)
	}, "#/components/parameters/")

	cs.PathParameters = NewMapRefSelf[PathParameter, *openapi3.ParameterRef](spec.Parameters, func(pr *openapi3.ParameterRef) (ref string, _ Ref[PathParameter]) {
		if pr.Ref != "" {
			return pr.Ref, nil
		}
		return "", NewPathParameter(pr.Value, cs.Schemas)
	}, "#/components/parameters/")

	cs.CookieParameters = NewMapRefSelf[CookieParameter, *openapi3.ParameterRef](spec.Parameters, func(pr *openapi3.ParameterRef) (ref string, _ Ref[CookieParameter]) {
		if pr.Ref != "" {
			return pr.Ref, nil
		}
		return "", NewCookieParameter(pr.Value, cs.Schemas)
	}, "#/components/parameters/")

	cs.Headers = NewMapRefSelf[Header, *openapi3.HeaderRef](spec.Headers, func(hr *openapi3.HeaderRef) (ref string, _ Ref[Header]) {
		if hr.Ref != "" {
			return hr.Ref, nil
		}
		return "", NewHeader(hr.Value, cs.Schemas)
	}, "#/components/headers/")

	cs.SecuritySchemes = NewMapRefSelf[SecurityScheme, *openapi3.SecuritySchemeRef](spec.SecuritySchemes, func(ss *openapi3.SecuritySchemeRef) (ref string, _ Ref[SecurityScheme]) {
		if ss.Ref != "" {
			return ss.Ref, nil
		}
		return "", NewSecurityScheme(ss.Value)
	}, "#/components/securitySchemes/")

	cs.Links = NewMapRefSelf[Link, *openapi3.LinkRef](spec.Links, func(lr *openapi3.LinkRef) (ref string, _ Ref[Link]) {
		if lr.Ref != "" {
			return lr.Ref, nil
		}
		return "", NewLink(lr.Value)
	}, "#/components/links/")

	return cs
}
