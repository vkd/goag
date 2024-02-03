package generator

import "github.com/vkd/goag/specification"

type Header struct {
	Spec specification.Header

	FieldName string
	Key       string
}

func NewHeader(spec specification.Header) Header {
	h := Header{
		Spec: spec,

		FieldName: PublicFieldName(spec.Name),
		Key:       spec.Name,
	}
	return h
}
