package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AuthMiddleware(header, token string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if r.Header.Get(header) != token {
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}
			h.ServeHTTP(rw, r)
		})
	}
}

func TestApiKeySecurity(t *testing.T) {
	testHeader := "Authorization"
	testToken := "1234"

	api := API{
		GetPetsHandler: func(_ context.Context, _ GetPetsRequest) GetPetsResponse { return NewGetPetsResponse200() },
	}
	api.Middlewares = append(api.Middlewares, AuthMiddleware(testHeader, "Bearer "+testToken))

	for _, tt := range []struct {
		token      string
		expectCode int
	}{
		{"Bearer " + testToken, 200},
		{"Bearer" + testToken, 401},
		{"" + testToken, 401},
		{"Bearer", 401},
		{"Bearer ", 401},
		{"Bearer 5678", 401},
	} {
		t.Run(tt.token, func(t *testing.T) {
			tt := tt

			req := httptest.NewRequest("GET", "/pets", nil)
			req.Header.Add(testHeader, tt.token)
			w := httptest.NewRecorder()
			api.ServeHTTP(w, req)
			assert.Equal(t, tt.expectCode, w.Code)
		})
	}
}

func TestApiKeySecurity_NotFound(t *testing.T) {
	testHeader := "Authorization"
	testToken := "1234"

	api := API{
		GetPetsHandler: func(_ context.Context, _ GetPetsRequest) GetPetsResponse { return NewGetPetsResponse200() },
	}
	api.Middlewares = append(api.Middlewares, func(h http.Handler) http.Handler {
		t.Fail()
		return h
	})

	req := httptest.NewRequest("GET", "/not_found", nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	req = httptest.NewRequest("GET", "/not_found", nil)
	req.Header.Add(testHeader, "Bearer "+testToken)
	w = httptest.NewRecorder()
	api.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestMiddleware_Order(t *testing.T) {
	var mx sync.Mutex
	var check []string
	flag := func(s string) {
		mx.Lock()
		check = append(check, s)
		mx.Unlock()
	}
	checkMiddleware := func(name string) func(http.Handler) http.Handler {
		return func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				flag(name + " in")
				h.ServeHTTP(rw, r)
				flag(name + " out")
			})
		}
	}

	api := API{
		GetPetsHandler: func(_ context.Context, _ GetPetsRequest) GetPetsResponse {
			flag("handler")
			return NewGetPetsResponse200()
		},
	}

	api.Middlewares = append(api.Middlewares,
		checkMiddleware("1"),
		checkMiddleware("2"),
		checkMiddleware("3"),
	)

	r := httptest.NewRequest("GET", "/pets", nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, strings.Split("1 in|2 in|3 in|handler|3 out|2 out|1 out", "|"), check)
}
