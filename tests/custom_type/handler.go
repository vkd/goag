package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/vkd/goag/tests/custom_type/pkg"
)

// ---------------------------------------------
// GetShopsShop -
// GET /shops/{shop}
// ---------------------------------------------

type GetShopsShopHandlerFunc func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse

func (f GetShopsShopHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetShopsShopHTTPRequest(r)).writeGetShopsShop(w)
}

func (GetShopsShopHandlerFunc) Path() string { return "/shops/{shop}" }

func (GetShopsShopHandlerFunc) Method() string { return http.MethodGet }

type GetShopsShopRequest interface {
	HTTP() *http.Request
	Parse() (GetShopsShopParams, error)
}

func GetShopsShopHTTPRequest(r *http.Request) GetShopsShopRequest {
	return getShopsShopHTTPRequest{r}
}

type getShopsShopHTTPRequest struct {
	Request *http.Request
}

func (r getShopsShopHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getShopsShopHTTPRequest) Parse() (GetShopsShopParams, error) {
	return newGetShopsShopParams(r.Request)
}

type GetShopsShopParams struct {
	Query GetShopsShopParamsQuery

	Path GetShopsShopParamsPath

	Body GetShop
}

type GetShopsShopParamsQuery struct {
	PageSchemaRefQuery pkg.Maybe[pkg.Page]

	PageCustomTypeQuery pkg.Maybe[pkg.PageCustomTypeQuery]
}

type GetShopsShopParamsPath struct {
	Shop pkg.Shop
}

func newGetShopsShopParams(r *http.Request) (zero GetShopsShopParams, _ error) {
	var params GetShopsShopParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["page_schema_ref_query"]
			if ok && len(q) > 0 {
				var vCustom string
				if len(q) == 1 {
					vCustom = q[0]
				} else {
					return zero, ErrParseParam{In: "query", Parameter: "page_schema_ref_query", Reason: "multiple values found: single value expected"}
				}
				var vOpt pkg.Page
				{
					err := vOpt.ParseString(vCustom)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "page_schema_ref_query", Reason: "parse custom type", Err: err}
					}
				}
				params.Query.PageSchemaRefQuery.Set(vOpt)
			}
		}
		{
			q, ok := query["page_custom_type_query"]
			if ok && len(q) > 0 {
				var vCustom string
				if len(q) == 1 {
					vCustom = q[0]
				} else {
					return zero, ErrParseParam{In: "query", Parameter: "page_custom_type_query", Reason: "multiple values found: single value expected"}
				}
				var vOpt pkg.PageCustomTypeQuery
				{
					err := vOpt.ParseString(vCustom)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "page_custom_type_query", Reason: "parse custom type", Err: err}
					}
				}
				params.Query.PageCustomTypeQuery.Set(vOpt)
			}
		}
	}

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}'")
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
				return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "required"}
			}

			vCustom := vPath
			{
				err := params.Path.Shop.ParseString(vCustom)
				if err != nil {
					return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "parse custom type", Err: err}
				}
			}
		}
	}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

func (r GetShopsShopParams) HTTP() *http.Request { return nil }

func (r GetShopsShopParams) Parse() (GetShopsShopParams, error) { return r, nil }

type GetShopsShopResponse interface {
	writeGetShopsShop(http.ResponseWriter)
}

func NewGetShopsShopResponse200JSON(body Shop) GetShopsShopResponse {
	var out GetShopsShopResponse200JSON
	out.Body = body
	return out
}

// GetShopsShopResponse200JSON - Shop response
type GetShopsShopResponse200JSON struct {
	Body Shop
}

func (r GetShopsShopResponse200JSON) writeGetShopsShop(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsShopResponse200JSON) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	writeJSON(w, r.Body, "GetShopsShopResponse200JSON")
}

func NewGetShopsShopResponseDefault(code int) GetShopsShopResponse {
	var out GetShopsShopResponseDefault
	out.Code = code
	return out
}

// GetShopsShopResponseDefault - Error response
type GetShopsShopResponseDefault struct {
	Code int
}

func (r GetShopsShopResponseDefault) writeGetShopsShop(w http.ResponseWriter) {
	r.Write(w)
}

func (r GetShopsShopResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

func writeJSON(w io.Writer, v interface{}, name string) {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		LogError(fmt.Errorf("write json response %q: %w", name, err))
	}
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
