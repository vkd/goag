package test

import (
	"fmt"
	"strconv"
)

type Page int

func (p Page) String() string { return strconv.Itoa(int(p)) }

func (p *Page) UnmarshalText(data []byte) error {
	i, err := strconv.Atoi(string(data))
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	*p = Page(i)
	return nil
}

type Shop string

func (s Shop) String() string { return string(s) }

func (s *Shop) UnmarshalText(data []byte) error {
	*s = Shop(string(data))
	return nil
}

type RequestID string

func (s RequestID) String() string { return string(s) }

func (s *RequestID) UnmarshalText(data []byte) error {
	*s = RequestID(string(data))
	return nil
}
