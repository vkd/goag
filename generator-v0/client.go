package generator

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/vkd/goag/generator-v0/source"
)

func NewClientBuilder(spec *openapi3.Swagger, handlers []Handler) (interface {
	Build() (*source.ClientFile, error)
}, error) {
	cb := &ClientBuilder{
		Hanlders: handlers,
	}
	// for _, h := range handlers {
	// 	for _, q := range h.Parameters.Queries {
	// 		cb.qu
	// 		q.Parameter.Name
	// 	}
	// }
	return cb, nil
}

type ClientBuilder struct {
	Hanlders []Handler
}

func (c *ClientBuilder) Build() (*source.ClientFile, error) {
	goFile := source.ClientFile{}

	for _, h := range c.Hanlders {
		f := &source.ClientFunc{
			Name:   h.Name,
			Method: h.Method,
			URL:    h.Path,
		}
		for _, q := range h.Params.Query {
			f.Queries = append(f.Queries, source.ClientFuncQuery{
				QueryName:        q.Parameter.Name,
				RequestFieldName: q.Field.Name,
				Formatter:        q.Field.Type.Format("request." + q.Field.Name),
			})
		}
		for _, r := range h.Responses {
			if r.IsDefault {
				f.DefaultResponse = source.ClientFuncResponse{
					ResponseType: r.PrivateName,
					HasBody:      r.IsBody,
				}
			} else {
				f.Responses = append(f.Responses, source.ClientFuncResponse{
					StatusCode:   r.Status,
					ResponseType: r.PrivateName,
					HasBody:      r.IsBody,
				})
			}
		}
		goFile.Funcs = append(goFile.Funcs, f)
	}
	return &goFile, nil
}
