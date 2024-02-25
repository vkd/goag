package specification

import "github.com/getkin/kin-openapi/openapi3"

type PathItem struct {
	NoRef[PathItem]

	// PathV2     Path // TODO
	Operations []*Operation

	// Deprecated // TODO
	Path     PathOld2
	PathOld  PathOld
	PathItem *openapi3.PathItem
}

func NewPathItem(path string) *PathItem {
	return &PathItem{}
}
