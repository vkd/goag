package test

import (
	"strconv"
)

type Page int

func (p Page) String() string { return strconv.Itoa(int(p)) }

func (p Page) Strings() []string { return []string{strconv.Itoa(int(p))} }

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
