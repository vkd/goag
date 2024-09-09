package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ---------------------------------------------
// PostShopsShopStringSepShopSchemaPets -
// ---------------------------------------------

type PostShopsShopStringSepShopSchemaPetsHandlerFunc func(ctx context.Context, r PostShopsShopStringSepShopSchemaPetsRequest) PostShopsShopStringSepShopSchemaPetsResponse

func (f PostShopsShopStringSepShopSchemaPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsShopStringSepShopSchemaPetsHTTPRequest(r)).writePostShopsShopStringSepShopSchemaPets(w)
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
	Query PostShopsShopStringSepShopSchemaPetsParamsQuery

	Path PostShopsShopStringSepShopSchemaPetsParamsPath

	Headers PostShopsShopStringSepShopSchemaPetsParamsHeaders
}

type PostShopsShopStringSepShopSchemaPetsParamsQuery struct {
	PageInt Maybe[int]

	PageSchema Maybe[Page]

	PagesSchema Maybe[Pages]

	PageIntReq int

	PageSchemaReq Page
}

type PostShopsShopStringSepShopSchemaPetsParamsPath struct {
	ShopString string

	ShopSchema Shop
}

type PostShopsShopStringSepShopSchemaPetsParamsHeaders struct {
	XOrganizationInt Maybe[int]

	XOrganizationSchema Maybe[Organization]

	XOrganizationIntRequired int

	XOrganizationSchemaRequired Organization
}

func newPostShopsShopStringSepShopSchemaPetsParams(r *http.Request) (zero PostShopsShopStringSepShopSchemaPetsParams, _ error) {
	var params PostShopsShopStringSepShopSchemaPetsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["page_int"]
			if ok && len(q) > 0 {
				vInt64, err := strconv.ParseInt(q[0], 10, 0)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_int", Reason: "parse int", Err: err}
				}
				v := int(vInt64)
				params.Query.PageInt.Set(v)
			}
		}
		{
			q, ok := query["page_schema"]
			if ok && len(q) > 0 {
				vInt64, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_schema", Reason: "parse int32", Err: err}
				}
				vInt32 := int32(vInt64)
				v := NewPage(vInt32)
				params.Query.PageSchema.Set(v)
			}
		}
		{
			q, ok := query["pages_schema"]
			if ok && len(q) > 0 {
				vInt32s := make([]int32, len(q))
				for i := range q {
					vInt64, err := strconv.ParseInt(q[i], 10, 32)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "pages_schema", Reason: "parse int32", Err: err}
					}
					vInt32s[i] = int32(vInt64)
				}
				v := NewPages(vInt32s)
				params.Query.PagesSchema.Set(v)
			}
		}
		{
			q, ok := query["page_int_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'page_int_req': is required")
			}
			if ok && len(q) > 0 {
				vInt64, err := strconv.ParseInt(q[0], 10, 0)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_int_req", Reason: "parse int", Err: err}
				}
				params.Query.PageIntReq = int(vInt64)
			}
		}
		{
			q, ok := query["page_schema_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'page_schema_req': is required")
			}
			if ok && len(q) > 0 {
				vInt64, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "page_schema_req", Reason: "parse int32", Err: err}
				}
				vInt32 := int32(vInt64)
				params.Query.PageSchemaReq = NewPage(vInt32)
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("X-Organization-Int")
			if len(hs) > 0 {
				vInt64, err := strconv.ParseInt(hs[0], 10, 0)
				if err != nil {
					return zero, ErrParseParam{In: "header", Parameter: "X-Organization-Int", Reason: "parse int", Err: err}
				}
				v := int(vInt64)
				params.Headers.XOrganizationInt.Set(v)
			}
		}
		{
			hs := header.Values("X-Organization-Schema")
			if len(hs) > 0 {
				vInt64, err := strconv.ParseInt(hs[0], 10, 0)
				if err != nil {
					return zero, ErrParseParam{In: "header", Parameter: "X-Organization-Schema", Reason: "parse int", Err: err}
				}
				vInt := int(vInt64)
				v := NewOrganization(vInt)
				params.Headers.XOrganizationSchema.Set(v)
			}
		}
		{
			hs := header.Values("X-Organization-Int-Required")
			if len(hs) == 0 {
				return zero, fmt.Errorf("header parameter 'X-Organization-Int-Required': is required")
			}
			if len(hs) > 0 {
				vInt64, err := strconv.ParseInt(hs[0], 10, 0)
				if err != nil {
					return zero, ErrParseParam{In: "header", Parameter: "X-Organization-Int-Required", Reason: "parse int", Err: err}
				}
				params.Headers.XOrganizationIntRequired = int(vInt64)
			}
		}
		{
			hs := header.Values("X-Organization-Schema-Required")
			if len(hs) == 0 {
				return zero, fmt.Errorf("header parameter 'X-Organization-Schema-Required': is required")
			}
			if len(hs) > 0 {
				vInt64, err := strconv.ParseInt(hs[0], 10, 0)
				if err != nil {
					return zero, ErrParseParam{In: "header", Parameter: "X-Organization-Schema-Required", Reason: "parse int", Err: err}
				}
				vInt := int(vInt64)
				params.Headers.XOrganizationSchemaRequired = NewOrganization(vInt)
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

			params.Path.ShopString = vPath
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

			vString := vPath
			vShopc := NewShopc(vString)
			vShopb := NewShopb(vShopc)
			vShopa := NewShopa(vShopb)
			params.Path.ShopSchema = NewShop(vShopa)
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
	writePostShopsShopStringSepShopSchemaPets(http.ResponseWriter)
}

func NewPostShopsShopStringSepShopSchemaPetsResponse200() PostShopsShopStringSepShopSchemaPetsResponse {
	var out PostShopsShopStringSepShopSchemaPetsResponse200
	return out
}

// PostShopsShopStringSepShopSchemaPetsResponse200 - OK response
type PostShopsShopStringSepShopSchemaPetsResponse200 struct{}

func (r PostShopsShopStringSepShopSchemaPetsResponse200) writePostShopsShopStringSepShopSchemaPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsShopStringSepShopSchemaPetsResponse200) Write(w http.ResponseWriter) {
	w.WriteHeader(200)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

type Maybe[T any] struct {
	IsSet bool
	Value T
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{}
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Maybe[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Maybe[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
}

type Nullable[T any] struct {
	IsSet bool
	Value T
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{}
}

func Pointer[T any](v T) Nullable[T] {
	return Nullable[T]{
		IsSet: true,
		Value: v,
	}
}

func (m Nullable[T]) Get() (zero T, _ bool) {
	if m.IsSet {
		return m.Value, true
	}
	return zero, false
}

func (m *Nullable[T]) Set(v T) {
	m.IsSet = true
	m.Value = v
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
