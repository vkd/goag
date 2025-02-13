package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/vkd/goag/tests/schema_one_of/pkg"
)

// ------------------------
//         Schemas
// ------------------------

type Cat struct {
	Label string
}

var _ json.Marshaler = (*Cat)(nil)

func (c Cat) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	write := func(bs []byte) {
		if err != nil {
			return
		}
		n, werr := out.Write(bs)
		if werr != nil {
			err = werr
		} else if len(bs) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}

	write([]byte(`{`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`}`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c Cat) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
	var err error
	var comma string
	write := func(s string) {
		if err != nil || len(s) == 0 {
			return
		}
		n, werr := out.Write([]byte(s))
		if werr != nil {
			err = werr
		} else if len(s) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}
	writeProperty := func(name string, v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(comma + `"` + name + `":null`)
		} else {
			write(comma + `"` + name + `":`)
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
		comma = ","
	}
	_ = writeProperty
	{
		var v any
		v = c.Label
		writeProperty("label", v)
	}

	return err
}

var _ json.Unmarshaler = (*Cat)(nil)

func (c *Cat) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Cat) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["label"]; ok {
		err = json.Unmarshal(raw, &c.Label)
		if err != nil {
			return fmt.Errorf("'label' field: unmarshal string: %w", err)
		}
		delete(m, "label")
	} else {
		return fmt.Errorf("'label' key is missing")
	}
	return nil
}

type Cat2 struct {
	Name    pkg.Maybe[string]
	PetType string
}

var _ json.Marshaler = (*Cat2)(nil)

func (c Cat2) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	write := func(bs []byte) {
		if err != nil {
			return
		}
		n, werr := out.Write(bs)
		if werr != nil {
			err = werr
		} else if len(bs) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}

	write([]byte(`{`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`}`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c Cat2) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
	var err error
	var comma string
	write := func(s string) {
		if err != nil || len(s) == 0 {
			return
		}
		n, werr := out.Write([]byte(s))
		if werr != nil {
			err = werr
		} else if len(s) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}
	writeProperty := func(name string, v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(comma + `"` + name + `":null`)
		} else {
			write(comma + `"` + name + `":`)
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
		comma = ","
	}
	_ = writeProperty
	if vOpt, ok := c.Name.Get(); ok {
		var v any = nil
		v = vOpt
		writeProperty("name", v)
	}
	{
		var v any
		v = c.PetType
		writeProperty("petType", v)
	}

	return err
}

var _ json.Unmarshaler = (*Cat2)(nil)

func (c *Cat2) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Cat2) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["name"]; ok {
		err = json.Unmarshal(raw, &c.Name.Value)
		if err != nil {
			return fmt.Errorf("'name' field: unmarshal string: %w", err)
		}
		c.Name.IsSet = true
		delete(m, "name")
	}
	if raw, ok := m["petType"]; ok {
		err = json.Unmarshal(raw, &c.PetType)
		if err != nil {
			return fmt.Errorf("'petType' field: unmarshal string: %w", err)
		}
		delete(m, "petType")
	} else {
		return fmt.Errorf("'petType' key is missing")
	}
	return nil
}

type Dog struct {
	Name string
	Tag  pkg.Maybe[string]
}

var _ json.Marshaler = (*Dog)(nil)

func (c Dog) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	write := func(bs []byte) {
		if err != nil {
			return
		}
		n, werr := out.Write(bs)
		if werr != nil {
			err = werr
		} else if len(bs) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}

	write([]byte(`{`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`}`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c Dog) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
	var err error
	var comma string
	write := func(s string) {
		if err != nil || len(s) == 0 {
			return
		}
		n, werr := out.Write([]byte(s))
		if werr != nil {
			err = werr
		} else if len(s) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}
	writeProperty := func(name string, v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(comma + `"` + name + `":null`)
		} else {
			write(comma + `"` + name + `":`)
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
		comma = ","
	}
	_ = writeProperty
	{
		var v any
		v = c.Name
		writeProperty("name", v)
	}
	if vOpt, ok := c.Tag.Get(); ok {
		var v any = nil
		v = vOpt
		writeProperty("tag", v)
	}

	return err
}

var _ json.Unmarshaler = (*Dog)(nil)

func (c *Dog) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Dog) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["name"]; ok {
		err = json.Unmarshal(raw, &c.Name)
		if err != nil {
			return fmt.Errorf("'name' field: unmarshal string: %w", err)
		}
		delete(m, "name")
	} else {
		return fmt.Errorf("'name' key is missing")
	}
	if raw, ok := m["tag"]; ok {
		err = json.Unmarshal(raw, &c.Tag.Value)
		if err != nil {
			return fmt.Errorf("'tag' field: unmarshal string: %w", err)
		}
		c.Tag.IsSet = true
		delete(m, "tag")
	}
	return nil
}

type Dog2 struct {
	Name    pkg.Maybe[string]
	PetType string
	Tag     pkg.Maybe[string]
}

var _ json.Marshaler = (*Dog2)(nil)

func (c Dog2) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	write := func(bs []byte) {
		if err != nil {
			return
		}
		n, werr := out.Write(bs)
		if werr != nil {
			err = werr
		} else if len(bs) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}

	write([]byte(`{`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`}`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c Dog2) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
	var err error
	var comma string
	write := func(s string) {
		if err != nil || len(s) == 0 {
			return
		}
		n, werr := out.Write([]byte(s))
		if werr != nil {
			err = werr
		} else if len(s) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}
	writeProperty := func(name string, v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(comma + `"` + name + `":null`)
		} else {
			write(comma + `"` + name + `":`)
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
		comma = ","
	}
	_ = writeProperty
	if vOpt, ok := c.Name.Get(); ok {
		var v any = nil
		v = vOpt
		writeProperty("name", v)
	}
	{
		var v any
		v = c.PetType
		writeProperty("petType", v)
	}
	if vOpt, ok := c.Tag.Get(); ok {
		var v any = nil
		v = vOpt
		writeProperty("tag", v)
	}

	return err
}

var _ json.Unmarshaler = (*Dog2)(nil)

func (c *Dog2) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Dog2) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["name"]; ok {
		err = json.Unmarshal(raw, &c.Name.Value)
		if err != nil {
			return fmt.Errorf("'name' field: unmarshal string: %w", err)
		}
		c.Name.IsSet = true
		delete(m, "name")
	}
	if raw, ok := m["petType"]; ok {
		err = json.Unmarshal(raw, &c.PetType)
		if err != nil {
			return fmt.Errorf("'petType' field: unmarshal string: %w", err)
		}
		delete(m, "petType")
	} else {
		return fmt.Errorf("'petType' key is missing")
	}
	if raw, ok := m["tag"]; ok {
		err = json.Unmarshal(raw, &c.Tag.Value)
		if err != nil {
			return fmt.Errorf("'tag' field: unmarshal string: %w", err)
		}
		c.Tag.IsSet = true
		delete(m, "tag")
	}
	return nil
}

type Pet struct {
	Cat    pkg.Maybe[Cat]
	Dog    pkg.Maybe[pkg.Dog]
	OneOf2 pkg.Maybe[PetOneOf2]
	OneOf3 pkg.Maybe[int64]
	OneOf4 pkg.Maybe[string]
}

func NewPetCat(v Cat) Pet {
	var out Pet
	out.Cat.Set(v)
	return out
}

func NewPetDog(v pkg.Dog) Pet {
	var out Pet
	out.Dog.Set(v)
	return out
}

func NewPetOneOf2(v PetOneOf2) Pet {
	var out Pet
	out.OneOf2.Set(v)
	return out
}

func NewPetOneOf3(v int64) Pet {
	var out Pet
	out.OneOf3.Set(v)
	return out
}

func NewPetOneOf4(v string) Pet {
	var out Pet
	out.OneOf4.Set(v)
	return out
}

var _ json.Marshaler = (*Pet)(nil)

func (c Pet) MarshalJSON() ([]byte, error) {
	if oneOfValue, ok := c.Cat.Get(); ok {
		bs, err := oneOfValue.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("oneOf [0] field: marshal object: %w", err)
		}
		return bs, nil
	}
	if oneOfValue, ok := c.Dog.Get(); ok {
		var v Dog
		var vc pkg.Dog
		vc = oneOfValue
		v = Dog(vc.ToSchemaDog())
		bs, err := v.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("oneOf [1] field: marshal object: %w", err)
		}
		return bs, nil
	}
	if oneOfValue, ok := c.OneOf2.Get(); ok {
		bs, err := oneOfValue.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("oneOf [2] field: marshal object: %w", err)
		}
		return bs, nil
	}
	if oneOfValue, ok := c.OneOf3.Get(); ok {
		bs, err := json.Marshal(&oneOfValue)
		if err != nil {
			return nil, fmt.Errorf("oneOf [3] field: marshal json for primitive type: %w", err)
		}
		return bs, nil
	}
	if oneOfValue, ok := c.OneOf4.Get(); ok {
		bs, err := json.Marshal(&oneOfValue)
		if err != nil {
			return nil, fmt.Errorf("oneOf [4] field: marshal json for primitive type: %w", err)
		}
		return bs, nil
	}

	return nil, fmt.Errorf("cannot marshal oneOf object: all field are empty")
}

var _ json.Unmarshaler = (*Pet)(nil)

func (c *Pet) UnmarshalJSON(bs []byte) error {
	var err error

	err = c.unmarshalJSON_Cat(bs)
	if err == nil {
		return nil
	}

	err = c.unmarshalJSON_Dog(bs)
	if err == nil {
		return nil
	}

	err = c.unmarshalJSON_OneOf2(bs)
	if err == nil {
		return nil
	}

	err = c.unmarshalJSON_OneOf3(bs)
	if err == nil {
		return nil
	}

	err = c.unmarshalJSON_OneOf4(bs)
	if err == nil {
		return nil
	}
	return fmt.Errorf("cannot unmarshal oneOf object: %w", err)
}

func (c *Pet) unmarshalJSON_Cat(bs []byte) error {
	var err error
	var oneOf0 Cat
	err = oneOf0.UnmarshalJSON(bs)
	if err != nil {
		return fmt.Errorf("oneOf [0] field: unmarshal ref type 'Cat': %w", err)
	}

	c.Cat.Set(oneOf0)
	return nil
}

func (c *Pet) unmarshalJSON_Dog(bs []byte) error {
	var err error
	var oneOf1 pkg.Dog

	var vRef Dog
	err = vRef.UnmarshalJSON(bs)
	if err != nil {
		return fmt.Errorf("oneOf [1] field: unmarshal ref type 'Dog': %w", err)
	}

	var cv pkg.Dog
	err = cv.SetFromSchemaDog(struct {
		Name string
		Tag  pkg.Maybe[string]
	}(vRef))
	if err != nil {
		return fmt.Errorf("oneOf [1] field: set from schema: %w", err)
	}
	oneOf1 = cv

	c.Dog.Set(oneOf1)
	return nil
}

func (c *Pet) unmarshalJSON_OneOf2(bs []byte) error {
	var err error
	var oneOf2 PetOneOf2
	err = oneOf2.UnmarshalJSON(bs)
	if err != nil {
		return fmt.Errorf("oneOf [2] field: unmarshal ref type 'PetOneOf2': %w", err)
	}

	c.OneOf2.Set(oneOf2)
	return nil
}

func (c *Pet) unmarshalJSON_OneOf3(bs []byte) error {
	var err error
	var v int64
	err = json.Unmarshal(bs, &v)
	if err != nil {
		return fmt.Errorf("oneOf [3] field: unmarshal int64: %w", err)
	}
	oneOf3 := v

	c.OneOf3.Set(oneOf3)
	return nil
}

func (c *Pet) unmarshalJSON_OneOf4(bs []byte) error {
	var err error
	var oneOf4 string
	err = json.Unmarshal(bs, &oneOf4)
	if err != nil {
		return fmt.Errorf("oneOf [4] field: unmarshal string: %w", err)
	}

	c.OneOf4.Set(oneOf4)
	return nil
}

type Pet2 struct {
	Cat2 pkg.Maybe[Cat2]
	Dog2 pkg.Maybe[pkg.Dog2]
}

func NewPet2Cat2(v Cat2) Pet2 {
	var out Pet2
	out.Cat2.Set(v)
	return out
}

func NewPet2Dog2(v pkg.Dog2) Pet2 {
	var out Pet2
	out.Dog2.Set(v)
	return out
}

var _ json.Marshaler = (*Pet2)(nil)

func (c Pet2) MarshalJSON() ([]byte, error) {
	if oneOfValue, ok := c.Cat2.Get(); ok {
		bs, err := oneOfValue.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("oneOf [0] field: marshal object: %w", err)
		}
		return bs, nil
	}
	if oneOfValue, ok := c.Dog2.Get(); ok {
		var v Dog2
		var vc pkg.Dog2
		vc = oneOfValue
		v = Dog2(vc.ToSchemaDog2())
		bs, err := v.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("oneOf [1] field: marshal object: %w", err)
		}
		return bs, nil
	}

	return nil, fmt.Errorf("cannot marshal oneOf object: all field are empty")
}

var _ json.Unmarshaler = (*Pet2)(nil)

func (c *Pet2) UnmarshalJSON(bs []byte) error {
	type tp struct {
		Key string `json:"petType"`
	}
	var v tp
	err := json.Unmarshal(bs, &v)
	if err != nil {
		return fmt.Errorf("cannot unmarshal discriminator: %w", err)
	}
	switch v.Key {
	case "Cat2", "cati":
		return c.unmarshalJSON_Cat2(bs)
	case "Dog2", "dogi":
		return c.unmarshalJSON_Dog2(bs)
	default:
		return fmt.Errorf("unknown discriminator: %q", v.Key)
	}
}

func (c *Pet2) unmarshalJSON_Cat2(bs []byte) error {
	var err error
	var oneOf0 Cat2
	err = oneOf0.UnmarshalJSON(bs)
	if err != nil {
		return fmt.Errorf("oneOf [0] field: unmarshal ref type 'Cat2': %w", err)
	}

	c.Cat2.Set(oneOf0)
	return nil
}

func (c *Pet2) unmarshalJSON_Dog2(bs []byte) error {
	var err error
	var oneOf1 pkg.Dog2

	var vRef Dog2
	err = vRef.UnmarshalJSON(bs)
	if err != nil {
		return fmt.Errorf("oneOf [1] field: unmarshal ref type 'Dog2': %w", err)
	}

	var cv pkg.Dog2
	err = cv.SetFromSchemaDog2(struct {
		Name    pkg.Maybe[string]
		PetType string
		Tag     pkg.Maybe[string]
	}(vRef))
	if err != nil {
		return fmt.Errorf("oneOf [1] field: set from schema: %w", err)
	}
	oneOf1 = cv

	c.Dog2.Set(oneOf1)
	return nil
}

type Resp struct {
	Pet Pet
}

var _ json.Marshaler = (*Resp)(nil)

func (c Resp) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	write := func(bs []byte) {
		if err != nil {
			return
		}
		n, werr := out.Write(bs)
		if werr != nil {
			err = werr
		} else if len(bs) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}

	write([]byte(`{`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`}`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c Resp) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
	var err error
	var comma string
	write := func(s string) {
		if err != nil || len(s) == 0 {
			return
		}
		n, werr := out.Write([]byte(s))
		if werr != nil {
			err = werr
		} else if len(s) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}
	writeProperty := func(name string, v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(comma + `"` + name + `":null`)
		} else {
			write(comma + `"` + name + `":`)
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
		comma = ","
	}
	_ = writeProperty
	{
		var v any
		v = c.Pet
		writeProperty("pet", v)
	}

	return err
}

var _ json.Unmarshaler = (*Resp)(nil)

func (c *Resp) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Resp) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["pet"]; ok {
		err = c.Pet.UnmarshalJSON(raw)
		if err != nil {
			return fmt.Errorf("'pet' field: unmarshal ref type 'Pet': %w", err)
		}
		delete(m, "pet")
	} else {
		return fmt.Errorf("'pet' key is missing")
	}
	return nil
}

type Resp2 struct {
	Pet Pet2
}

var _ json.Marshaler = (*Resp2)(nil)

func (c Resp2) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	write := func(bs []byte) {
		if err != nil {
			return
		}
		n, werr := out.Write(bs)
		if werr != nil {
			err = werr
		} else if len(bs) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}

	write([]byte(`{`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`}`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c Resp2) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
	var err error
	var comma string
	write := func(s string) {
		if err != nil || len(s) == 0 {
			return
		}
		n, werr := out.Write([]byte(s))
		if werr != nil {
			err = werr
		} else if len(s) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}
	writeProperty := func(name string, v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(comma + `"` + name + `":null`)
		} else {
			write(comma + `"` + name + `":`)
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
		comma = ","
	}
	_ = writeProperty
	{
		var v any
		v = c.Pet
		writeProperty("pet", v)
	}

	return err
}

var _ json.Unmarshaler = (*Resp2)(nil)

func (c *Resp2) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *Resp2) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["pet"]; ok {
		err = c.Pet.UnmarshalJSON(raw)
		if err != nil {
			return fmt.Errorf("'pet' field: unmarshal ref type 'Pet2': %w", err)
		}
		delete(m, "pet")
	} else {
		return fmt.Errorf("'pet' key is missing")
	}
	return nil
}

type PetOneOf2 struct {
	ID int64
}

var _ json.Marshaler = (*PetOneOf2)(nil)

func (c PetOneOf2) MarshalJSON() ([]byte, error) {
	var out bytes.Buffer
	var err error
	write := func(bs []byte) {
		if err != nil {
			return
		}
		n, werr := out.Write(bs)
		if werr != nil {
			err = werr
		} else if len(bs) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}

	write([]byte(`{`))
	mErr := c.marshalJSONInnerBody(&out)
	if mErr != nil {
		err = mErr
	}
	write([]byte(`}`))

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (c PetOneOf2) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
	var err error
	var comma string
	write := func(s string) {
		if err != nil || len(s) == 0 {
			return
		}
		n, werr := out.Write([]byte(s))
		if werr != nil {
			err = werr
		} else if len(s) != n {
			err = fmt.Errorf("wrong len of written body")
		}
	}
	writeProperty := func(name string, v any) {
		if err != nil {
			return
		}
		if v == nil {
			write(comma + `"` + name + `":null`)
		} else {
			write(comma + `"` + name + `":`)
			werr := encoder.Encode(v)
			if werr != nil {
				err = werr
			}
		}
		comma = ","
	}
	_ = writeProperty
	{
		var v any
		v = c.ID
		writeProperty("id", v)
	}

	return err
}

var _ json.Unmarshaler = (*PetOneOf2)(nil)

func (c *PetOneOf2) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	return c.unmarshalJSONInnerBody(m)
}

func (c *PetOneOf2) unmarshalJSONInnerBody(m map[string]json.RawMessage) error {
	var err error
	_ = err
	if raw, ok := m["id"]; ok {
		var v int64
		err = json.Unmarshal(raw, &v)
		if err != nil {
			return fmt.Errorf("'id' field: unmarshal int64: %w", err)
		}
		c.ID = v
		delete(m, "id")
	} else {
		return fmt.Errorf("'id' key is missing")
	}
	return nil
}
