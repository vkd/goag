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

		Int *int

		Int32Req int32

		Int32 *int32

		Int64Req int64

		Int64 *int64

		Float32Req float32

		Float32 *float32

		Float64Req float64

		Float64 *float64

		StringReq string

		String *string

		Tag []string

		Filter []int32
	}

	Path struct {
		Shop int32
	}

	Headers struct {
		RequestID *string

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
				v := int(vInt)
				params.Query.IntReq = v
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
				params.Query.Int = &v
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
				v := int32(vInt)
				params.Query.Int32Req = v
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
				params.Query.Int32 = &v
			}
		}
		{
			q, ok := query["int64_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'int64_req': is required")
			}
			if ok && len(q) > 0 {
				v, err := strconv.ParseInt(q[0], 10, 64)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "int64_req", Reason: "parse int64", Err: err}
				}
				params.Query.Int64Req = v
			}
		}
		{
			q, ok := query["int64"]
			if ok && len(q) > 0 {
				v, err := strconv.ParseInt(q[0], 10, 64)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "int64", Reason: "parse int64", Err: err}
				}
				params.Query.Int64 = &v
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
				v := float32(vFloat)
				params.Query.Float32Req = v
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
				params.Query.Float32 = &v
			}
		}
		{
			q, ok := query["float64_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'float64_req': is required")
			}
			if ok && len(q) > 0 {
				v, err := strconv.ParseFloat(q[0], 64)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "float64_req", Reason: "parse float64", Err: err}
				}
				params.Query.Float64Req = v
			}
		}
		{
			q, ok := query["float64"]
			if ok && len(q) > 0 {
				v, err := strconv.ParseFloat(q[0], 64)
				if err != nil {
					return zero, ErrParseParam{In: "query", Parameter: "float64", Reason: "parse float64", Err: err}
				}
				params.Query.Float64 = &v
			}
		}
		{
			q, ok := query["string_req"]
			if !ok {
				return zero, fmt.Errorf("query parameter 'string_req': is required")
			}
			if ok && len(q) > 0 {
				v := q[0]
				params.Query.StringReq = v
			}
		}
		{
			q, ok := query["string"]
			if ok && len(q) > 0 {
				v := q[0]
				params.Query.String = &v
			}
		}
		{
			q, ok := query["tag"]
			if ok && len(q) > 0 {
				params.Query.Tag = q
			}
		}
		{
			q, ok := query["filter"]
			if ok && len(q) > 0 {
				params.Query.Filter = make([]int32, len(q))
				for i := range q {
					vInt, err := strconv.ParseInt(q[i], 10, 32)
					if err != nil {
						return zero, ErrParseParam{In: "query", Parameter: "filter", Reason: "parse int32", Err: err}
					}
					v1 := int32(vInt)
					params.Query.Filter[i] = v1
				}
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
				params.Headers.RequestID = &v
			}
		}
		{
			hs := header.Values("user-id")
			if len(hs) == 0 {
				return zero, fmt.Errorf("header parameter 'user-id': is required")
			}
			if len(hs) > 0 {
				v := hs[0]
				params.Headers.UserID = v
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
			v := int32(vInt)
			params.Path.Shop = v
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