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
// GetReviews -
// ---------------------------------------------

type GetReviewsHandlerFunc func(ctx context.Context, r GetReviewsRequest) GetReviewsResponse

func (f GetReviewsHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(r.Context(), GetReviewsHTTPRequest(r)).Write(w)
}

type GetReviewsRequest interface {
	HTTP() *http.Request
	Parse() (GetReviewsParams, error)
}

func GetReviewsHTTPRequest(r *http.Request) GetReviewsRequest {
	return getReviewsHTTPRequest{r}
}

type getReviewsHTTPRequest struct {
	Request *http.Request
}

func (r getReviewsHTTPRequest) HTTP() *http.Request { return r.Request }

func (r getReviewsHTTPRequest) Parse() (GetReviewsParams, error) {
	return newGetReviewsParams(r.Request)
}

type GetReviewsParams struct {
	Query struct {
		IntReq int

		Int Maybe[int]

		Int32Req int32

		Int32 Maybe[int32]

		Int64Req int64

		Int64 Maybe[int64]

		Float32Req float32

		Float32 Maybe[float32]

		Float64Req float64

		Float64 Maybe[float64]

		StringReq string

		String Maybe[string]

		Tag Maybe[[]string]

		Filter Maybe[[]int32]
	}

	Path struct {
		Shop int32
	}

	Headers struct {
		RequestID Maybe[string]

		UserID string
	}
}

func newGetReviewsParams(r *http.Request) (zero GetReviewsParams, _ error) {
	var params GetReviewsParams

	// Query parameters
	{
		query := r.URL.Query()
		{
			q, ok := query["int_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'int_req': is required")
			}
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 0)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "int_req", Reason: "parse int", Err: err}
				}
				params.Query.IntReq = int(vInt)
			}
		}
		{
			q, ok := query["int"]
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 0)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "int", Reason: "parse int", Err: err}
				}
				v := int(vInt)
				params.Query.Int.Set(v)
			}
		}
		{
			q, ok := query["int32_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'int32_req': is required")
			}
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "int32_req", Reason: "parse int32", Err: err}
				}
				params.Query.Int32Req = int32(vInt)
			}
		}
		{
			q, ok := query["int32"]
			if ok && len(q) > 0 {
				vInt, err := strconv.ParseInt(q[0], 10, 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "int32", Reason: "parse int32", Err: err}
				}
				v := int32(vInt)
				params.Query.Int32.Set(v)
			}
		}
		{
			q, ok := query["int64_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'int64_req': is required")
			}
			if ok && len(q) > 0 {
				var err error
				params.Query.Int64Req, err = strconv.ParseInt(q[0], 10, 64)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "int64_req", Reason: "parse int64", Err: err}
				}
			}
		}
		{
			q, ok := query["int64"]
			if ok && len(q) > 0 {
				v, err := strconv.ParseInt(q[0], 10, 64)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "int64", Reason: "parse int64", Err: err}
				}
				params.Query.Int64.Set(v)
			}
		}
		{
			q, ok := query["float32_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'float32_req': is required")
			}
			if ok && len(q) > 0 {
				vFloat, err := strconv.ParseFloat(q[0], 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "float32_req", Reason: "parse float32", Err: err}
				}
				params.Query.Float32Req = float32(vFloat)
			}
		}
		{
			q, ok := query["float32"]
			if ok && len(q) > 0 {
				vFloat, err := strconv.ParseFloat(q[0], 32)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "float32", Reason: "parse float32", Err: err}
				}
				v := float32(vFloat)
				params.Query.Float32.Set(v)
			}
		}
		{
			q, ok := query["float64_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'float64_req': is required")
			}
			if ok && len(q) > 0 {
				var err error
				params.Query.Float64Req, err = strconv.ParseFloat(q[0], 64)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "float64_req", Reason: "parse float64", Err: err}
				}
			}
		}
		{
			q, ok := query["float64"]
			if ok && len(q) > 0 {
				v, err := strconv.ParseFloat(q[0], 64)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "float64", Reason: "parse float64", Err: err}
				}
				params.Query.Float64.Set(v)
			}
		}
		{
			q, ok := query["string_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'string_req': is required")
			}
			if ok && len(q) > 0 {
				params.Query.StringReq = q[0]
			}
		}
		{
			q, ok := query["string"]
			if ok && len(q) > 0 {
				v := q[0]
				params.Query.String.Set(v)
			}
		}
		{
			q, ok := query["tag"]
			if ok && len(q) > 0 {
				v := q
				params.Query.Tag.Set(v)
			}
		}
		{
			q, ok := query["filter"]
			if ok && len(q) > 0 {
				v := make([]int32, len(q))
				for i := range q {
					vInt, err := strconv.ParseInt(q[i], 10, 32)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "filter", Reason: "parse int32", Err: err}
					}
					v[i] = int32(vInt)
				}
				params.Query.Filter.Set(v)
			}
		}
	}

	// Headers
	{
		header := r.Header
		{
			hs := header.Values("request-id")
			if len(hs) > 0 {
				v := hs[0]
				params.Headers.RequestID.Set(v)
			}
		}
		{
			hs := header.Values("user-id")
			if len(hs) == 0 {
				return zero, fmt.Errorf("header parameter 'user-id': is required")
			}
			if len(hs) > 0 {
				params.Headers.UserID = hs[0]
			}
		}
	}

	// Path parameters
	{
		p := r.URL.Path

		if !strings.HasPrefix(p, "/shops/") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/reviews'")
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

			vInt, err := strconv.ParseInt(vPath, 10, 32)
			if err != nil {
				return zero, ErrParseParam{In: "path", Parameter: "shop", Reason: "parse int32", Err: err}
			}
			params.Path.Shop = int32(vInt)
		}

		if !strings.HasPrefix(p, "/reviews") {
			return zero, fmt.Errorf("wrong path: expected '/shops/{shop}/reviews'")
		}
		p = p[8:] // "/reviews"
	}

	return params, nil
}

func (r GetReviewsParams) HTTP() *http.Request { return nil }

func (r GetReviewsParams) Parse() (GetReviewsParams, error) { return r, nil }

type GetReviewsResponse interface {
	getReviews()
	Write(w http.ResponseWriter)
}

func NewGetReviewsResponseDefault(code int) GetReviewsResponse {
	var out GetReviewsResponseDefault
	out.Code = code
	return out
}

type GetReviewsResponseDefault struct {
	Code int
}

func (r GetReviewsResponseDefault) getReviews() {}

func (r GetReviewsResponseDefault) Write(w http.ResponseWriter) {
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
		IsSet: true,
		Value: v,
	}
}

func (m *Maybe[T]) Set(v T) {
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
