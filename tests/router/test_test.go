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
		GetRTHandler: GetRTHandlerFunc(func(_ GetRTRequester) GetRTResponder { return GetRTResponseDefault(201) }),

		GetShopsHandler: GetShopsHandlerFunc(func(_ GetShopsRequester) GetShopsResponder { return GetShopsResponseDefault(202) }),

		GetShopsRTHandler: GetShopsRTHandlerFunc(func(_ GetShopsRTRequester) GetShopsRTResponder { return GetShopsRTResponseDefault(203) }),

		GetShopsShopHandler: GetShopsShopHandlerFunc(func(r GetShopsShopRequester) GetShopsShopResponder {
			_, err := r.Parse()
			if err != nil {
				return GetShopsShopResponseDefault(400)
			}
			return GetShopsShopResponseDefault(204)
		}),

		GetShopsShopRTHandler: GetShopsShopRTHandlerFunc(func(r GetShopsShopRTRequester) GetShopsShopRTResponder {
			_, err := r.Parse()
			if err != nil {
				return GetShopsShopRTResponseDefault(400)
			}
			return GetShopsShopRTResponseDefault(205)
		}),

		GetShopsShopPetsHandler: GetShopsShopPetsHandlerFunc(func(r GetShopsShopPetsRequester) GetShopsShopPetsResponder {
			_, err := r.Parse()
			if err != nil {
				return GetShopsShopPetsResponseDefault(400)
			}
			return GetShopsShopPetsResponseDefault(206)
		}),

		GetShopsActivateHandler: GetShopsActivateHandlerFunc(func(_ GetShopsActivateRequester) GetShopsActivateResponder {
			return GetShopsActivateResponseDefault(207)
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
			api.Middlewares = append(api.Middlewares, func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
					path, _ := SchemaPath(r)
					assert.Equal(t, tt.schemaPath, path)
					h.ServeHTTP(rw, r)
				})
			})
			w := httptest.NewRecorder()
			path := "/api/v1" + tt.path
			api.ServeHTTP(w, httptest.NewRequest("GET", path, nil))
			assert.Equal(t, tt.code, w.Code, "path: %s", tt.path)
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
