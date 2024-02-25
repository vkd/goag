package generator

import (
	"strings"

	"github.com/vkd/goag/specification"
)

type Response struct {
	s *specification.ResponseOld

	Name       string
	StatusCode string

	Headers []Header

	Body Optional[any]
}

func NewResponse(handlerName OperationName, response *specification.ResponseOld) *Response {
	r := &Response{s: response}
	r.Name = string(handlerName) + "Response" + strings.Title(response.StatusCode)
	if len(response.Spec.Content) > 0 {
		r.Name += "JSON"
		r.Body.OK = true
	}
	r.StatusCode = response.StatusCode

	for _, header := range response.Headers {
		r.Headers = append(r.Headers, NewHeader(header))
	}
	return r
}
