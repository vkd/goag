package test

import (
	_ "embed"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	api := API{
		GetHandler: GetHandlerFunc(func(_ GetRequest) GetResponse { return NewGetResponseDefault(201) }),

		GetShopsHandler: GetShopsHandlerFunc(func(_ GetShopsRequest) GetShopsResponse { return NewGetShopsResponseDefault(202) }),

		GetShopsRTHandler: GetShopsRTHandlerFunc(func(_ GetShopsRTRequest) GetShopsRTResponse { return NewGetShopsRTResponseDefault(203) }),

		GetShopsShopHandler: GetShopsShopHandlerFunc(func(r GetShopsShopRequest) GetShopsShopResponse {
			_, err := r.Parse()
			if err != nil {
				return NewGetShopsShopResponseDefault(400)
			}
			return NewGetShopsShopResponseDefault(204)
		}),

		GetShopsShopRTHandler: GetShopsShopRTHandlerFunc(func(r GetShopsShopRTRequest) GetShopsShopRTResponse {
			_, err := r.Parse()
			if err != nil {
				return NewGetShopsShopRTResponseDefault(400)
			}
			return NewGetShopsShopRTResponseDefault(205)
		}),

		GetShopsShopPetsHandler: GetShopsShopPetsHandlerFunc(func(r GetShopsShopPetsRequest) GetShopsShopPetsResponse {
			_, err := r.Parse()
			if err != nil {
				return NewGetShopsShopPetsResponseDefault(400)
			}
			return NewGetShopsShopPetsResponseDefault(206)
		}),

		GetShopsActivateHandler: GetShopsActivateHandlerFunc(func(_ GetShopsActivateRequest) GetShopsActivateResponse {
			return NewGetShopsActivateResponseDefault(207)
		}),
	}

	for _, tt := range []struct {
		path       string
		code       int
		schemaPath string
	}{
		{"/", 201, "/"},
		{"/shops", 202, "/shops"},
		{"/shops/", 203, "/shops/"},
		{"/shops/my_shop", 204, "/shops/{shop}"},

		{"/shops/my_shop/", 205, "/shops/{shop}/"},
		{"/shops/1/", 205, "/shops/{shop}/"},

		{"/shops/my_shop/pets", 206, "/shops/{shop}/pets"},

		{"/shops/activate", 207, "/shops/activate"},

		{"/not_found", 404, "/this_is_not_gonna_be_checked"},
	} {
		tt := tt
		t.Run(tt.path, func(t *testing.T) {
			api := api
			var mCount int
			api.Middlewares = append(api.Middlewares, func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
					path, _ := SchemaPath(r)
					assert.Equal(t, tt.schemaPath, path)
					h.ServeHTTP(rw, r)
					mCount++
				})
			})
			w := httptest.NewRecorder()
			path := "/api/v1" + tt.path
			api.ServeHTTP(w, httptest.NewRequest("GET", path, nil))
			assert.Equal(t, tt.code, w.Code, "path: %s", tt.path)
			switch tt.code {
			case http.StatusNotFound:
				assert.Equal(t, 0, mCount)
			default:
				assert.Equal(t, 1, mCount)
			}
		})
	}
}

//go:embed openapi.yaml
var openapiSpec string

func TestRounter_SpecFile(t *testing.T) {
	assert.Equal(t, openapiSpec, SpecFile)
}

func TestRouter_SpecHandle_NotFound(t *testing.T) {
	// not found if SpecFileHandler is nil
	api := API{}
	r := httptest.NewRequest("GET", "/api/v1/openapi.yaml", nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestRouter_SpecHandle(t *testing.T) {
	api := API{
		SpecFileHandler: SpecFileHandler(),
	}
	r := httptest.NewRequest("GET", "/api/v1/openapi.yaml", nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, SpecFile, w.Body.String())
}

func TestRouter_SpecHandle_with_Middleware(t *testing.T) {
	// serve spec despite of middlewares
	api := API{
		SpecFileHandler: SpecFileHandler(),
	}
	api.Middlewares = append(api.Middlewares, func(_ http.Handler) http.Handler {
		return http.NotFoundHandler()
	})
	r := httptest.NewRequest("GET", "/api/v1/openapi.yaml", nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, SpecFile, w.Body.String())
}
