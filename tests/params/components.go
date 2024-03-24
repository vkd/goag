package test

import (
	"fmt"
	"strconv"
)

// ---------------------------------
//         Query Parameters
// ---------------------------------

type Float32 float32

func (q *Float32) ParseQuery(vs []string) error {
	var v float32
	vFloat, err := strconv.ParseFloat(vs[0], 32)
	if err != nil {
		return fmt.Errorf("parse float32: %w", err)
	}
	v = float32(vFloat)
	*q = Float32(v)
	return nil
}

func (q Float32) String() string {
	return strconv.FormatFloat(float64(float32(q)), 'e', -1, 32)
}

type Float32Required float32

func (q *Float32Required) ParseQuery(vs []string) error {
	var v float32
	vFloat, err := strconv.ParseFloat(vs[0], 32)
	if err != nil {
		return fmt.Errorf("parse float32: %w", err)
	}
	v = float32(vFloat)
	*q = Float32Required(v)
	return nil
}

func (q Float32Required) String() string {
	return strconv.FormatFloat(float64(float32(q)), 'e', -1, 32)
}

type Float64 float64

func (q *Float64) ParseQuery(vs []string) error {
	var v float64
	var err error
	v, err = strconv.ParseFloat(vs[0], 64)
	if err != nil {
		return fmt.Errorf("parse float64: %w", err)
	}
	*q = Float64(v)
	return nil
}

func (q Float64) String() string {
	return strconv.FormatFloat(float64(q), 'e', -1, 64)
}

type Float64Required float64

func (q *Float64Required) ParseQuery(vs []string) error {
	var v float64
	var err error
	v, err = strconv.ParseFloat(vs[0], 64)
	if err != nil {
		return fmt.Errorf("parse float64: %w", err)
	}
	*q = Float64Required(v)
	return nil
}

func (q Float64Required) String() string {
	return strconv.FormatFloat(float64(q), 'e', -1, 64)
}

type Int int

func (q *Int) ParseQuery(vs []string) error {
	var v int
	vInt, err := strconv.ParseInt(vs[0], 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v = int(vInt)
	*q = Int(v)
	return nil
}

func (q Int) String() string {
	return strconv.FormatInt(int64(int(q)), 10)
}

type Int32 int32

func (q *Int32) ParseQuery(vs []string) error {
	var v int32
	vInt, err := strconv.ParseInt(vs[0], 10, 32)
	if err != nil {
		return fmt.Errorf("parse int32: %w", err)
	}
	v = int32(vInt)
	*q = Int32(v)
	return nil
}

func (q Int32) String() string {
	return strconv.FormatInt(int64(int32(q)), 10)
}

type Int32Required int32

func (q *Int32Required) ParseQuery(vs []string) error {
	var v int32
	vInt, err := strconv.ParseInt(vs[0], 10, 32)
	if err != nil {
		return fmt.Errorf("parse int32: %w", err)
	}
	v = int32(vInt)
	*q = Int32Required(v)
	return nil
}

func (q Int32Required) String() string {
	return strconv.FormatInt(int64(int32(q)), 10)
}

type Int64 int64

func (q *Int64) ParseQuery(vs []string) error {
	var v int64
	var err error
	v, err = strconv.ParseInt(vs[0], 10, 64)
	if err != nil {
		return fmt.Errorf("parse int64: %w", err)
	}
	*q = Int64(v)
	return nil
}

func (q Int64) String() string {
	return strconv.FormatInt(int64(q), 10)
}

type Int64Required int64

func (q *Int64Required) ParseQuery(vs []string) error {
	var v int64
	var err error
	v, err = strconv.ParseInt(vs[0], 10, 64)
	if err != nil {
		return fmt.Errorf("parse int64: %w", err)
	}
	*q = Int64Required(v)
	return nil
}

func (q Int64Required) String() string {
	return strconv.FormatInt(int64(q), 10)
}

type IntRequired int

func (q *IntRequired) ParseQuery(vs []string) error {
	var v int
	vInt, err := strconv.ParseInt(vs[0], 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v = int(vInt)
	*q = IntRequired(v)
	return nil
}

func (q IntRequired) String() string {
	return strconv.FormatInt(int64(int(q)), 10)
}

type String string

func (q *String) ParseQuery(vs []string) error {
	var v string
	v = vs[0]
	*q = String(v)
	return nil
}

func (q String) String() string {
	return string(q)
}

type StringRequired string

func (q *StringRequired) ParseQuery(vs []string) error {
	var v string
	v = vs[0]
	*q = StringRequired(v)
	return nil
}

func (q StringRequired) String() string {
	return string(q)
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
