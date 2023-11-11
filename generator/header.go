package generator

import "github.com/vkd/goag/specification"

type Header struct {
	Spec     specification.Header
	Response *Response

	FieldName string
	Key       string
}

func NewHeader(r *Response, spec specification.Header) Header {
	h := Header{
		Spec:     spec,
		Response: r,

		FieldName: PublicFieldName(spec.Name),
		Key:       spec.Name,
	}
	return h
}
