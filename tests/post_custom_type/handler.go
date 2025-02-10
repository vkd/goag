package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	pkg "github.com/vkd/goag/tests/post_custom_type/pkg1"
)

// ---------------------------------------------
// PostShopsShopPets -
// ---------------------------------------------

type PostShopsShopPetsHandlerFunc func(ctx context.Context, r PostShopsShopPetsRequest) PostShopsShopPetsResponse

func (f PostShopsShopPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsShopPetsHTTPRequest(r)).writePostShopsShopPets(w)
}

type PostShopsShopPetsRequest interface {
	HTTP() *http.Request
	Parse() (PostShopsShopPetsParams, error)
}

func PostShopsShopPetsHTTPRequest(r *http.Request) PostShopsShopPetsRequest {
	return postShopsShopPetsHTTPRequest{r}
}

type postShopsShopPetsHTTPRequest struct {
	Request *http.Request
}

func (r postShopsShopPetsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r postShopsShopPetsHTTPRequest) Parse() (PostShopsShopPetsParams, error) {
	return newPostShopsShopPetsParams(r.Request)
}

type PostShopsShopPetsParams struct {
	Query PostShopsShopPetsParamsQuery

	Path PostShopsShopPetsParamsPath

	Body NewPet
}

type PostShopsShopPetsParamsQuery struct {
	Filter pkg.Maybe[pkg.ShopType]
}

type PostShopsShopPetsParamsPath struct {
	Shop pkg.ShopType
}

func newPostShopsShopPetsParams(r *http.Request) (zero PostShopsShopPetsParams, _ error) {
	var params PostShopsShopPetsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["filter"]
			if ok && len(q) > 0 {
				vCustom := q[0]
				var vOpt pkg.ShopType
				{
					err := vOpt.ParseString(vCustom)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "filter", Reason: "parse custom type", Err: err}
					}
				}
				params.Query.Filter.Set(vOpt)
			}
		}
	}

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets'")
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

		if !strings.HasPrefix(p, "/pets") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/pets'")
		}
		p = p[5:] // "/pets"
	}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&params.Body)
	if err != nil {
		return zero, fmt.Errorf("decode request body: %w", err)
	}

	return params, nil
}

func (r PostShopsShopPetsParams) HTTP() *http.Request { return nil }

func (r PostShopsShopPetsParams) Parse() (PostShopsShopPetsParams, error) { return r, nil }

type PostShopsShopPetsResponse interface {
	writePostShopsShopPets(http.ResponseWriter)
}

func NewPostShopsShopPetsResponse201() PostShopsShopPetsResponse {
	var out PostShopsShopPetsResponse201
	return out
}

// PostShopsShopPetsResponse201 - OK response
type PostShopsShopPetsResponse201 struct{}

func (r PostShopsShopPetsResponse201) writePostShopsShopPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsShopPetsResponse201) Write(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func NewPostShopsShopPetsResponseDefault(code int) PostShopsShopPetsResponse {
	var out PostShopsShopPetsResponseDefault
	out.Code = code
	return out
}

// PostShopsShopPetsResponseDefault - Default response
type PostShopsShopPetsResponseDefault struct {
	Code int
}

func (r PostShopsShopPetsResponseDefault) writePostShopsShopPets(w http.ResponseWriter) {
	r.Write(w)
}

func (r PostShopsShopPetsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

func write(w io.Writer, r io.Reader, name string) {
	_, err := io.Copy(w, r)
	if err != nil {
		LogError(fmt.Errorf("write response %q: %w", name, err))
	}
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

var _ json.Marshaler = (*Nullable[any])(nil)

func (m Nullable[T]) MarshalJSON() ([]byte, error) {
	if m.IsSet {
		return json.Marshal(&m.Value)
	}
	return []byte(nullValue), nil
}

var _ json.Unmarshaler = (*Nullable[any])(nil)

const nullValue = "null"

var nullValueBs = []byte(nullValue)

func (m *Nullable[T]) UnmarshalJSON(bs []byte) error {
	if bytes.Equal(bs, nullValueBs) {
		m.IsSet = false
		return nil
	}
	m.IsSet = true
	return json.Unmarshal(bs, &m.Value)
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
