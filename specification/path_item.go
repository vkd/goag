package specification

import "github.com/getkin/kin-openapi/openapi3"

type PathItem struct {
	NoRef[PathItem]

	RawPath    string // TODO
	Operations []*Operation

	// Deprecated // TODO
	Path     PathOld2
	PathItem *openapi3.PathItem
}

func NewPathItem(path string) *PathItem {
	return &PathItem{
		RawPath: path,
	}
}
