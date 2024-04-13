package test

import (
	"fmt"
	"strconv"
)

type Page int

func (p Page) String() string { return strconv.Itoa(int(p)) }

func (p Page) Strings() []string { return []string{strconv.Itoa(int(p))} }

func (p *Page) Parse(v string) error {
	i, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	*p = Page(i)
	return nil
}

type Shop string

func (s Shop) String() string { return string(s) }

func (s *Shop) Parse(v string) error {
	*s = Shop(v)
	return nil
}

type RequestID string

func (s RequestID) String() string { return string(s) }

func (s *RequestID) Parse(v string) error {
	*s = RequestID(v)
	return nil
}
