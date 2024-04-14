package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/vkd/goag/tests/post_custom_type/pkg"
)

// ---------------------------------------------
// PostShopsShopPets -
// ---------------------------------------------

type PostShopsShopPetsHandlerFunc func(ctx context.Context, r PostShopsShopPetsRequest) PostShopsShopPetsResponse

func (f PostShopsShopPetsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), PostShopsShopPetsHTTPRequest(r)).Write(w)
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
	Path struct {
		Shop pkg.ShopType
	}

	Body NewPet
}

func newPostShopsShopPetsParams(r *http.Request) (zero PostShopsShopPetsParams, _ error) {
	var params PostShopsShopPetsParams

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

			err := params.Path.Shop.Parse(vPath)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "parse custom type", Err: err}
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
	postShopsShopPets()
	Write(w http.ResponseWriter)
}

func NewPostShopsShopPetsResponse201() PostShopsShopPetsResponse {
	var out PostShopsShopPetsResponse201
	return out
}

type PostShopsShopPetsResponse201 struct{}

func (r PostShopsShopPetsResponse201) postShopsShopPets() {}

func (r PostShopsShopPetsResponse201) Write(w http.ResponseWriter) {
	w.WriteHeader(201)
}

func NewPostShopsShopPetsResponseDefault(code int) PostShopsShopPetsResponse {
	var out PostShopsShopPetsResponseDefault
	out.Code = code
	return out
}

type PostShopsShopPetsResponseDefault struct {
	Code int
}

func (r PostShopsShopPetsResponseDefault) postShopsShopPets() {}

func (r PostShopsShopPetsResponseDefault) Write(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
}

var LogError = func(err error) {
	log.Println(fmt.Sprintf("Error: %v", err))
}

type Maybe[T any] struct {
	IsSet bool
	Value T
}

func Just[T any](v T) Maybe[T] {
	return Maybe[T]{
		Value: v,
		IsSet: true,
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
