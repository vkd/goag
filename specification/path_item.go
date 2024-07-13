package specification

type PathItem struct {
	NoRef[PathItem]

	RawPath    string
	Operations []*Operation
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
