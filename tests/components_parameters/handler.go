package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------
// PostShopsShopStringSepShopSchemaPets -
// ---------------------------------------------

type PostShopsShopStringSepShopSchemaPetsHandlerFunc func(ctx context.Context, r PostShopsShopStringSepShopSchemaPetsRequest) PostShopsShopStringSepShopSchemaPetsResponse

func (f PostShopsShopStringSepShopSchemaPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsShopStringSepShopSchemaPetsHTTPRequest(r)).Write(w)
}

type PostShopsShopStringSepShopSchemaPetsRequest interface {
	HTTP() *http.Request
	Parse() (PostShopsShopStringSepShopSchemaPetsParams, error)
}

func PostShopsShopStringSepShopSchemaPetsHTTPRequest(r *http.Request) PostShopsShopStringSepShopSchemaPetsRequest {
	return postShopsShopStringSepShopSchemaPetsHTTPRequest{r}
}

type postShopsShopStringSepShopSchemaPetsHTTPRequest struct {
	Request *http.Request
}

func (r postShopsShopStringSepShopSchemaPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postShopsShopStringSepShopSchemaPetsHTTPRequest) Parse() (PostShopsShopStringSepShopSchemaPetsParams, error) {
	return newPostShopsShopStringSepShopSchemaPetsParams(r.Request)
}

type PostShopsShopStringSepShopSchemaPetsParams struct {
	Query struct {
		PageInt *PageIntQuery

		PageSchema *PageSchemaQuery

		PageIntReq PageIntQueryRequired

		PageSchemaReq PageSchemaQueryRequired
	}

	Path struct {
		ShopString ShopStringPath

		ShopSchema ShopSchemaPath
	}

	Headers struct {
		XOrganizationInt *OrgIntHeader

		XOrganizationSchema *OrgSchemaHeader

		XOrganizationIntRequired OrgIntHeaderRequired

		XOrganizationSchemaRequired OrgSchemaHeaderRequired
	}
}

func newPostShopsShopStringSepShopSchemaPetsParams(r *http.Request) (zero PostShopsShopStringSepShopSchemaPetsParams, _ error) {
	var params PostShopsShopStringSepShopSchemaPetsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["page_int"]
			if ok && len(q) > 0 {
				var v PageIntQuery
				err := v.ParseQuery(q)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_int", Reason: "parse PageIntQuery", Err: err}
				}
				params.Query.PageInt = &v
			}
		}
		{
			q, ok := query["page_schema"]
			if ok && len(q) > 0 {
				var v PageSchemaQuery
				err := v.ParseQuery(q)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_schema", Reason: "parse PageSchemaQuery", Err: err}
				}
				params.Query.PageSchema = &v
			}
		}
		{
			q, ok := query["page_int_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'page_int_req': is required")
			}
			if ok && len(q) > 0 {
				err := params.Query.PageIntReq.ParseQuery(q)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_int_req", Reason: "parse PageIntQueryRequired", Err: err}
				}
			}
		}
		{
			q, ok := query["page_schema_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'page_schema_req': is required")
			}
			if ok && len(q) > 0 {
				err := params.Query.PageSchemaReq.ParseQuery(q)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_schema_req", Reason: "parse PageSchemaQueryRequired", Err: err}
				}
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("X-Organization-Int")
			if len(hs) > 0 {
				var v OrgIntHeader
				err := v.Parse(hs[0])
				if err != nil {
					return zero, ErrParseParam{In: "header", Parameter: "X-Organization-Int", Reason: "parse OrgIntHeader", Err: err}
				}
				params.Headers.XOrganizationInt = &v
			}
		}
		{
			hs := header.Values("X-Organization-Schema")
			if len(hs) > 0 {
				var v OrgSchemaHeader
				err := v.Parse(hs[0])
				if err != nil {
					return zero, ErrParseParam{In: "header", Parameter: "X-Organization-Schema", Reason: "parse OrgSchemaHeader", Err: err}
				}
				params.Headers.XOrganizationSchema = &v
			}
		}
		{
			hs := header.Values("X-Organization-Int-Required")
			if len(hs) == 0 {
				return zero, fmt.Errorf("header parameter 'X-Organization-Int-Required': is required")
			}
			if len(hs) > 0 {
				err := params.Headers.XOrganizationIntRequired.Parse(hs[0])
				if err != nil {
					return zero, ErrParseParam{In: "header", Parameter: "X-Organization-Int-Required", Reason: "parse OrgIntHeaderRequired", Err: err}
				}
			}
		}
		{
			hs := header.Values("X-Organization-Schema-Required")
			if len(hs) == 0 {
				return zero, fmt.Errorf("header parameter 'X-Organization-Schema-Required': is required")
			}
			if len(hs) > 0 {
				err := params.Headers.XOrganizationSchemaRequired.Parse(hs[0])
				if err != nil {
					return zero, ErrParseParam{In: "header", Parameter: "X-Organization-Schema-Required", Reason: "parse OrgSchemaHeaderRequired", Err: err}
				}
			}
		}
	}

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop_string}/sep/{shop_schema}/pets'")
		}
		p = p[7:] // "/shops/"

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			if len(vPath) == 0 {
				return zero, ErrParseParam{In: "path", Parameter: "shop_string", Reason: "required"}
			}

			err := params.Path.ShopString.Parse(vPath)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "shop_string", Reason: "parse ShopStringPath", Err: err}
			}
		}

		if !strings.HasPrefix(p, "/sep/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop_string}/sep/{shop_schema}/pets'")
		}
		p = p[5:] // "/sep/"

		{
			idx := strings.Index(p, "/")
			if idx == -1 {
				idx = len(p)
			}
			vPath := p[:idx]
			p = p[idx:]

			if len(vPath) == 0 {
				return zero, ErrParseParam{In: "path", Parameter: "shop_schema", Reason: "required"}
			}

			err := params.Path.ShopSchema.Parse(vPath)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "shop_schema", Reason: "parse ShopSchemaPath", Err: err}
			}
		}

		if !strings.HasPrefix(p, "/pets") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop_string}/sep/{shop_schema}/pets'")
		}
		p = p[5:] // "/pets"
	}

	return params, nil
}

func (r PostShopsShopStringSepShopSchemaPetsParams) HTTP() *http.Request { return nil }

func (r PostShopsShopStringSepShopSchemaPetsParams) Parse() (PostShopsShopStringSepShopSchemaPetsParams, error) {
	return r, nil
}

type PostShopsShopStringSepShopSchemaPetsResponse interface {
	postShopsShopStringSepShopSchemaPets()
	Write(w http.ResponseWriter)
}

func NewPostShopsShopStringSepShopSchemaPetsResponse200() PostShopsShopStringSepShopSchemaPetsResponse {
	var out PostShopsShopStringSepShopSchemaPetsResponse200
	return out
}

// PostShopsShopStringSepShopSchemaPetsResponse200 - OK response
type PostShopsShopStringSepShopSchemaPetsResponse200 struct{}

func (r PostShopsShopStringSepShopSchemaPetsResponse200) postShopsShopStringSepShopSchemaPets() {}

func (r PostShopsShopStringSepShopSchemaPetsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

type ErrParseParam struct {
	In        string
	Parameter string
	Reason    string
	Err       error
}

func (e ErrParseParam) Error() string {
	return fmt.Sprintf("%s parameter '%s': %s: %v", e.In, e.Parameter, e.Reason, e.Err)
}

func (e ErrParseParam) Unwrap() error { return e.Err }
