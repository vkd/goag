package test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseDefault(t *testing.T) {
	handler := GetPetsHandlerFunc(func(_ GetPetsRequester) GetPetsResponder {
		return GetPetsResponseDefaultJSON(400, Error{Message: "test default response"})
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"message\":\"test default response\"}\n", w.Body.String())
}

type ResponseStatusWriter struct {
	rCtx context.Context
	http.ResponseWriter
	status    int
	statusOut chan<- int
}

func (r *ResponseStatusWriter) WriteHeader(s int) {
	select {
	case <-r.rCtx.Done():
		s = 499
	default:
	}
	r.status = s
	r.statusOut <- s
	r.ResponseWriter.WriteHeader(s)
}

func TestResponse_Canceled(t *testing.T) {
	clientStart := make(chan struct{})

	api := &API{
		GetPetsHandler: func(r GetPetsRequester) GetPetsResponder {
			req := r.Parse()
			ctx := req.HTTPRequest.Context()

			close(clientStart)
			<-ctx.Done()
			return GetPetsResponseDefaultJSON(500, Error{Message: "test default response"})
		},
	}
	status := make(chan int, 1)
	api.Middlewares = append(api.Middlewares, func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(&ResponseStatusWriter{
				statusOut:      status,
				ResponseWriter: rw,
				rCtx:           r.Context(),
			}, r)
		})
	})

	srv := httptest.NewServer(api)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-clientStart
		cancel()
	}()

	req, err := http.NewRequestWithContext(ctx, "GET", srv.URL+"/pets", nil)
	if err != nil {
		t.Errorf("Error on create request: %v", err)
	}

	client := srv.Client()
	_, err = client.Do(req)
	assert.True(t, errors.Is(err, context.Canceled))

	assert.Equal(t, 499, <-status)
}
