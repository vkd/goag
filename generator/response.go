package generator

import (
	"strings"

	"github.com/vkd/goag/specification"
)

type Response struct {
	Spec          specification.Response
	ClientHandler *ClientHandler
	Operation     *Operation

	PublicTypeName string
	StatusCode     string
	Description    string
	Headers        []Header

	// temporary
	HasBody bool
}

func NewResponse(ch *ClientHandler, o *Operation, spec specification.Response) Response {
	r := Response{
		Spec:          spec,
		ClientHandler: ch,
		Operation:     o,

		StatusCode: spec.StatusCode,
		HasBody:    len(spec.Spec.Content) > 0,
	}
	r.PublicTypeName = o.Name + "Response" + strings.Title(spec.StatusCode)
	if r.HasBody {
		r.PublicTypeName += "JSON"
	}
	for _, h := range spec.Headers {
		r.Headers = append(r.Headers, NewHeader(&r, h))
	}
	return r
}
