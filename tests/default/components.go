package test

import (
	"fmt"
	"strconv"
)

// ------------------------
//         Schemas
// ------------------------

type Error struct {
	Message string `json:"message"`
}

type NewPet struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type Pet struct {
	NewPet
	ID int64 `json:"id"`
}

type Pets []Pet

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

type PageSizeQuery int32

func (q *PageSizeQuery) ParseQuery(vs []string) error {
	var v int32
	vInt, err := strconv.ParseInt(vs[0], 10, 32)
	if err != nil {
		return fmt.Errorf("parse int32: %w", err)
	}
	v = int32(vInt)
	*q = PageSizeQuery(v)
	return nil
}

func (q PageSizeQuery) String() string {
	return strconv.FormatInt(int64(int32(q)), 10)
}

// --------------------------------
//         Path Parameters
// --------------------------------

type ShopPathRequired int32

func (q *ShopPathRequired) Parse(s string) error {
	var v int32
	vInt, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return fmt.Errorf("parse int32: %w", err)
	}
	v = int32(vInt)
	*q = ShopPathRequired(v)
	return nil
}

func (q ShopPathRequired) String() string {
	return strconv.FormatInt(int64(int32(q)), 10)
}
