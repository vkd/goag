package test

import (
	"fmt"
	"strconv"
)

// ------------------------
//         Schemas
// ------------------------

type Organization int

func (c *Organization) ParseString(s string) error {
	vInt, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v := int(vInt)
	*c = Organization(v)
	return nil
}

func (q Organization) String() string {
	return strconv.FormatInt(int64(int(q)), 10)
}

func (q Organization) Strings() []string {
	return []string{q.String()}
}

type Page int32

func (c *Page) ParseString(s string) error {
	vInt, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return fmt.Errorf("parse int32: %w", err)
	}
	v := int32(vInt)
	*c = Page(v)
	return nil
}

func (q Page) String() string {
	return strconv.FormatInt(int64(int32(q)), 10)
}

func (q Page) Strings() []string {
	return []string{q.String()}
}

type Pages []int32

func (c *Pages) ParseStrings(s []string) error {
	v := make([]int32, len(s))
	for i := range s {
		vInt, err := strconv.ParseInt(s[i], 10, 32)
		if err != nil {
			return fmt.Errorf("parse int32: %w", err)
		}
		v[i] = int32(vInt)
	}
	*c = Pages(v)
	return nil
}

func (q Pages) Strings() []string {
	out := make([]string, 0, len(q))
	for _, v := range q {
		out = append(out, strconv.FormatInt(int64(v), 10))
	}
	return out
}

type Shop Shopa

func (c *Shop) ParseString(s string) error {
	var v Shopa
	err := v.ParseString(s)
	if err != nil {
		return fmt.Errorf("parse Shopa: %w", err)
	}
	*c = Shop(v)
	return nil
}

func (q Shop) String() string {
	return Shopa(q).String()
}

func (q Shop) Strings() []string {
	return []string{q.String()}
}

type Shopa string

func (c *Shopa) ParseString(s string) error {
	v := s
	*c = Shopa(v)
	return nil
}

func (q Shopa) String() string {
	return string(q)
}

func (q Shopa) Strings() []string {
	return []string{q.String()}
}

type Shops []string

func (c *Shops) ParseStrings(s []string) error {
	v := s
	*c = Shops(v)
	return nil
}

func (q Shops) Strings() []string {
	out := make([]string, 0, len(q))
	for _, v := range q {
		out = append(out, v)
	}
	return out
}
