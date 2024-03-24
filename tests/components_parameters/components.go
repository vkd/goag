package test

import (
	"fmt"
	"strconv"
)

// ------------------------
//         Schemas
// ------------------------

type Organization int

func (c *Organization) Parse(s string) error {
	var v int
	vInt, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v = int(vInt)
	*c = Organization(v)
	return nil
}

func (q Organization) String() string {
	return strconv.FormatInt(int64(int(q)), 10)
}

type Page int32

func (c *Page) Parse(s string) error {
	var v int32
	vInt, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return fmt.Errorf("parse int32: %w", err)
	}
	v = int32(vInt)
	*c = Page(v)
	return nil
}

func (q Page) String() string {
	return strconv.FormatInt(int64(int32(q)), 10)
}

type Shop string

func (c *Shop) Parse(s string) error {
	var v string
	v = s
	*c = Shop(v)
	return nil
}

func (q Shop) String() string {
	return string(q)
}

// ---------------------------------
//         Query Parameters
// ---------------------------------

type PageIntQuery int

func (q *PageIntQuery) ParseQuery(vs []string) error {
	var v int
	vInt, err := strconv.ParseInt(vs[0], 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v = int(vInt)
	*q = PageIntQuery(v)
	return nil
}

func (q PageIntQuery) String() string {
	return strconv.FormatInt(int64(int(q)), 10)
}

type PageIntQueryRequired int

func (q *PageIntQueryRequired) ParseQuery(vs []string) error {
	var v int
	vInt, err := strconv.ParseInt(vs[0], 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v = int(vInt)
	*q = PageIntQueryRequired(v)
	return nil
}

func (q PageIntQueryRequired) String() string {
	return strconv.FormatInt(int64(int(q)), 10)
}

type PageSchemaQuery Page

func (q *PageSchemaQuery) ParseQuery(vs []string) error {
	var v Page
	err := v.Parse(vs[0])
	if err != nil {
		return fmt.Errorf("parse Page: %w", err)
	}
	*q = PageSchemaQuery(v)
	return nil
}

func (q PageSchemaQuery) String() string {
	return Page(q).String()
}

type PageSchemaQueryRequired Page

func (q *PageSchemaQueryRequired) ParseQuery(vs []string) error {
	var v Page
	err := v.Parse(vs[0])
	if err != nil {
		return fmt.Errorf("parse Page: %w", err)
	}
	*q = PageSchemaQueryRequired(v)
	return nil
}

func (q PageSchemaQueryRequired) String() string {
	return Page(q).String()
}

// ----------------------------------
//         Header Parameters
// ----------------------------------

type OrgIntHeader int

func (h *OrgIntHeader) Parse(s string) error {
	var v int
	vInt, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v = int(vInt)
	*h = OrgIntHeader(v)
	return nil
}

func (h OrgIntHeader) String() string {
	return strconv.FormatInt(int64(int(h)), 10)
}

type OrgIntHeaderRequired int

func (h *OrgIntHeaderRequired) Parse(s string) error {
	var v int
	vInt, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v = int(vInt)
	*h = OrgIntHeaderRequired(v)
	return nil
}

func (h OrgIntHeaderRequired) String() string {
	return strconv.FormatInt(int64(int(h)), 10)
}

type OrgSchemaHeader Organization

func (h *OrgSchemaHeader) Parse(s string) error {
	var v Organization
	err := v.Parse(s)
	if err != nil {
		return fmt.Errorf("parse Organization: %w", err)
	}
	*h = OrgSchemaHeader(v)
	return nil
}

func (h OrgSchemaHeader) String() string {
	return Organization(h).String()
}

type OrgSchemaHeaderRequired Organization

func (h *OrgSchemaHeaderRequired) Parse(s string) error {
	var v Organization
	err := v.Parse(s)
	if err != nil {
		return fmt.Errorf("parse Organization: %w", err)
	}
	*h = OrgSchemaHeaderRequired(v)
	return nil
}

func (h OrgSchemaHeaderRequired) String() string {
	return Organization(h).String()
}

// --------------------------------
//         Path Parameters
// --------------------------------

type ShopSchemaPath Shop

func (q *ShopSchemaPath) Parse(s string) error {
	var v Shop
	err := v.Parse(s)
	if err != nil {
		return fmt.Errorf("parse Shop: %w", err)
	}
	*q = ShopSchemaPath(v)
	return nil
}

func (q ShopSchemaPath) String() string {
	return Shop(q).String()
}

type ShopStringPath string

func (q *ShopStringPath) Parse(s string) error {
	var v string
	v = s
	*q = ShopStringPath(v)
	return nil
}

func (q ShopStringPath) String() string {
	return string(q)
}
