package pkg

type Page string

func (s Page) String() string { return string(s) }

func (s *Page) Parse(v string) error {
	*s = Page(v)
	return nil
}
