package test

import (
	"fmt"
)

type Page int

func (p Page) Int32() int32 { return int32(p) }

func (p *Page) ParseInt32(v int32) error {
	*p = Page(v)
	return nil
}

type Shop string

func (s Shop) String() string { return string(s) }

func (s *Shop) ParseString(v string) error {
	*s = Shop(v)
	return nil
}

type RequestID string

func (s RequestID) String() string { return string(s) }

func (s *RequestID) ParseString(v string) error {
	*s = RequestID(v)
	return nil
}

type Pages []Page

func (p *Pages) ParseInt32s(vs []int32) error {
	*p = make(Pages, 0, len(vs))
	for _, v := range vs {
		var page Page
		err := page.ParseInt32(v)
		if err != nil {
			return fmt.Errorf("parse Page for '%v' value: %w", v, err)
		}
		*p = append(*p, page)
	}
	return nil
}

func (p Pages) Int32s() []int32 {
	out := make([]int32, 0, len(p))
	for _, page := range p {
		out = append(out, page.Int32())
	}
	return out
}
