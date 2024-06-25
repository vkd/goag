package pkg

type Page string

func (s Page) String() string    { return string(s) }
func (s Page) Strings() []string { return []string{s.String()} }

func (s *Page) ParseString(v string) error {
	*s = Page(v)
	return nil
}
func (s *Page) Parse(v string) error {
	return s.ParseString(v)
}

type PageCustomTypeQuery string

func (s PageCustomTypeQuery) String() string    { return string(s) }
func (s PageCustomTypeQuery) Strings() []string { return []string{s.String()} }

func (s *PageCustomTypeQuery) ParseString(v string) error {
	*s = PageCustomTypeQuery(v)
	return nil
}
