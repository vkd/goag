package pkg

type Page string

func (s Page) String() string { return string(s) }

func (s *Page) Parse(v string) error {
	*s = Page(v)
	return nil
}

type PageCustomTypeQuery string

func (s PageCustomTypeQuery) String() string    { return string(s) }
func (s PageCustomTypeQuery) Strings() []string { return []string{s.String()} }

func (s *PageCustomTypeQuery) Parse(v string) error {
	*s = PageCustomTypeQuery(v)
	return nil
}
