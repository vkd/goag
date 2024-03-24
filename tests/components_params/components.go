package test

import (
	"fmt"
	"strconv"
)

// ---------------------------------
//         Query Parameters
// ---------------------------------

type PageQuery int32

func (q *PageQuery) ParseQuery(vs []string) error {
	var v int32
	vInt, err := strconv.ParseInt(vs[0], 10, 32)
	if err != nil {
		return fmt.Errorf("parse int32: %w", err)
	}
	v = int32(vInt)
	*q = PageQuery(v)
	return nil
}

func (q PageQuery) String() string {
	return strconv.FormatInt(int64(int32(q)), 10)
}

// --------------------------------
//         Path Parameters
// --------------------------------

type ShopPathRequired string

func (q *ShopPathRequired) Parse(s string) error {
	var v string
	v = s
	*q = ShopPathRequired(v)
	return nil
}

func (q ShopPathRequired) String() string {
	return string(q)
}
