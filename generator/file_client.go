package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Client struct {
	Imports []Import

	Operations []ClientOperation

	IsDecodeJSONFunc bool
}

func NewClient(s *specification.Spec, ops []*Operation) Client {
	c := Client{}
	c.Operations = make([]ClientOperation, 0, len(ops))
	for _, o := range ops {
		c.Operations = append(c.Operations, NewClientOperation(o))
	}
	return c
}

func (c Client) Render() (string, error) {
	return ExecuteTemplate("Client", c)
}

type ClientOperation struct {
	*Operation

	RequestVarName string
	IsRequestBody  bool

	Headers []ClientHeader
}

func NewClientOperation(o *Operation) ClientOperation {
	c := ClientOperation{
		Operation:     o,
		IsRequestBody: o.Operation.RequestBody.IsSet && o.Operation.RequestBody.Value.Value().Content.Has("application/json"),
	}

	c.Headers = make([]ClientHeader, 0, len(o.Params.Headers.List))
	for _, h := range o.Params.Headers.List {
		c.Headers = append(c.Headers, ClientHeader{
			Name:      h.V.Name,
			FieldName: h.V.FieldName,
			Required:  h.V.Required,
			Type:      h.V.Type,
		})
	}

	if len(o.Security) > 0 {
		for _, sr := range o.Security {
			if sr.Scheme.Type != specification.SecuritySchemeTypeHTTP {
				continue
			}
			if sr.Scheme.Scheme != "bearer" {
				continue
			}
			c.Headers = append(c.Headers, ClientHeader{
				Name:      "Authorization",
				FieldName: "Authorization",
				Required:  len(o.Security) == 1,
			})
		}
	}

	return c
}

func (c ClientOperation) PathFormat() (Renders, error) {
	var out Renders

	var v string
	for _, dir := range c.Operation.Operation.Path.Dirs {
		if !dir.IsVariable {
			v += "/" + dir.V
			continue
		}
		out = append(out, QuotedRender(v+"/"))
		v = ""

		param := dir.Param.Value()
		gp, ok := c.Operation.Params.Path.Get(param.Name)
		if !ok {
			return nil, fmt.Errorf("%q path parameter: not found in %q operation", param.Name, c.Operation.Name)
		}
		out = append(out, RenderFunc(func() (string, error) {
			return gp.V.Type.RenderFormat("request.Path." + gp.V.FieldName)
		}))
	}
	if v != "" {
		out = append(out, QuotedRender(v))
	}
	return out, nil
}

type ClientHeader struct {
	Name      string
	FieldName string
	Required  bool
	Type      Formatter
}

func (c *ClientHeader) RenderFormat(from string) (string, error) {
	switch t := c.Type.(type) {
	case CustomType:
		return t.RenderFormat(from)
	}
	if c.Required {
		return from, nil
	}
	return "*" + from, nil
}
