package pkg

import (
	"encoding/json"
	"fmt"
)

type Page string

func (s Page) String() string { return string(s) }

func (s *Page) ParseString(v string) error {
	*s = Page(v)
	return nil
}

type PageCustomTypeQuery string

func (s PageCustomTypeQuery) String() string { return string(s) }

func (s *PageCustomTypeQuery) ParseString(v string) error {
	*s = PageCustomTypeQuery(v)
	return nil
}

type Shop string

func (s *Shop) ParseString(str string) error {
	*s = Shop(str)
	return nil
}

func (s Shop) String() string {
	return string(s)
}

type Metadata struct {
	InternalID string
}

var _ json.Unmarshaler = (*Metadata)(nil)

func (m *Metadata) UnmarshalJSON(bs []byte) error {
	type tp Metadata
	var v tp
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return fmt.Errorf("unmarshal metadata: %w", err)
	}
	*m = Metadata(v)
	return nil
}
