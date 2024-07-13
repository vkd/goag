package specification

import "github.com/getkin/kin-openapi/openapi3"

type PathItem struct {
	NoRef[PathItem]

	RawPath    string // TODO
	Operations []*Operation

	// Deprecated // TODO
	PathItem *openapi3.PathItem
}

func NewPathItem(path string) *PathItem {
	return &PathItem{
		RawPath: path,
	}
}

func (p *PathItem) GetOperation(m HTTPMethod) (*Operation, bool) {
	for _, o := range p.Operations {
		if o.HTTPMethod == m {
			return o, true
		}
	}
	return nil, false
}

func (p *PathItem) HasOperation(m HTTPMethod) bool {
	_, ok := p.GetOperation(m)
	return ok
}
