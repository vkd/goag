package specification

import "github.com/getkin/kin-openapi/openapi3"

type QueryParameter struct {
	NoRef[QueryParameter]

	Name            string
	Description     string
	Required        bool
	Deprecated      bool
	AllowEmptyValue bool

	Style         string
	Explode       Maybe[bool]
	AllowReserved bool

	Schema Ref[Schema]
}

func NewQueryParameter(p *openapi3.Parameter, schemas ComponentsSchemas, opts SchemaOptions) *QueryParameter {
	return &QueryParameter{
		Name:            p.Name,
		Description:     p.Description,
		Required:        p.Required,
		Deprecated:      p.Deprecated,
		AllowEmptyValue: p.AllowEmptyValue,

		Style:         p.Style,
		Explode:       NewMaybe(p.Explode),
		AllowReserved: p.AllowReserved,

		Schema: NewSchemaRef(p.Schema, schemas, opts),
	}
}

var _ Ref[QueryParameter] = (*QueryParameter)(nil)

func (q *QueryParameter) Value() *QueryParameter { return q }

type HeaderParameter struct {
	NoRef[HeaderParameter]

	Name        string
	Description string
	Required    bool
	Deprecated  bool

	Style   string
	Explode Maybe[bool]

	Schema Ref[Schema]
}

func NewHeaderParameter(p *openapi3.Parameter, schemas ComponentsSchemas, opts SchemaOptions) *HeaderParameter {
	return &HeaderParameter{
		Name:        p.Name,
		Description: p.Description,
		Required:    p.Required,
		Deprecated:  p.Deprecated,

		Style:   p.Style,
		Explode: NewMaybe(p.Explode),

		Schema: NewSchemaRef(p.Schema, schemas, opts),
	}
}

var _ Ref[HeaderParameter] = (*HeaderParameter)(nil)

func (h *HeaderParameter) Value() *HeaderParameter { return h }

type PathParameter struct {
	NoRef[PathParameter]

	Name        string
	Description string
	Deprecated  bool

	Style   string
	Explode Maybe[bool]

	Schema Ref[Schema]
}

func NewPathParameter(p *openapi3.Parameter, schemas ComponentsSchemas, opts SchemaOptions) *PathParameter {
	return &PathParameter{
		Name:        p.Name,
		Description: p.Description,
		Deprecated:  p.Deprecated,

		Style:   p.Style,
		Explode: NewMaybe(p.Explode),

		Schema: NewSchemaRef(p.Schema, schemas, opts),
	}
}

var _ Ref[PathParameter] = (*PathParameter)(nil)

func (p *PathParameter) Value() *PathParameter { return p }

type CookieParameter struct {
	NoRef[CookieParameter]

	Name        string
	Description string
	Required    bool
	Deprecated  bool

	Style   string
	Explode Maybe[bool]

	Schema Ref[Schema]
}

func NewCookieParameter(p *openapi3.Parameter, schemas ComponentsSchemas, opts SchemaOptions) *CookieParameter {
	return &CookieParameter{
		Name:        p.Name,
		Description: p.Description,
		Required:    p.Required,
		Deprecated:  p.Deprecated,

		Style:   p.Style,
		Explode: NewMaybe(p.Explode),

		Schema: NewSchemaRef(p.Schema, schemas, opts),
	}
}

var _ Ref[CookieParameter] = (*CookieParameter)(nil)

func (c *CookieParameter) Value() *CookieParameter { return c }
