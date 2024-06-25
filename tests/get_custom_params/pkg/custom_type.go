package pkg

type Page string

func (s Page) String() string { return string(s) }

func (s *Page) ParseString(v string) error {
	*s = Page(v)
	return nil
}
