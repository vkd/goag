package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// ------------------------
//         Schemas
// ------------------------

type Error struct {
	Detail string `json:"detail"`
}

var _ json.Marshaler = (*Error)(nil)

func (c Error) MarshalJSON() ([]byte, error) {
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

func (c Error) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
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
	writeJSON := func(v any) {
		if err != nil {
			return
		}
		werr := encoder.Encode(v)
		if werr != nil {
			err = werr
		}
	}
	_ = writeJSON
	write([]byte(`"detail":`))
	writeJSON(c.Detail)

	return err
}

type Pet struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

var _ json.Marshaler = (*Pet)(nil)

func (c Pet) MarshalJSON() ([]byte, error) {
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

func (c Pet) marshalJSONInnerBody(out io.Writer) error {
	encoder := json.NewEncoder(out)
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
	writeJSON := func(v any) {
		if err != nil {
			return
		}
		werr := encoder.Encode(v)
		if werr != nil {
			err = werr
		}
	}
	_ = writeJSON
	write([]byte(`"id":`))
	writeJSON(c.ID)
	write([]byte(`,`))
	write([]byte(`"name":`))
	writeJSON(c.Name)

	return err
}

type Pets []Pet

// ------------------------
//         Headers
// ------------------------

// ErrorCode - Error code
type ErrorCode int

func (h ErrorCode) String() string {
	return strconv.FormatInt(int64(int(h)), 10)
}

func (c *ErrorCode) Parse(s string) error {
	vInt, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	v := int(vInt)
	*c = ErrorCode(v)
	return nil
}

// ------------------------------
//         RequestBodies
// ------------------------------

type CreatePetJSON struct {
	Name string `json:"name"`
}

// ------------------------------
//         Responses
// ------------------------------

func NewErrorResponseResponse(body Error, xErrorCode Maybe[int]) ErrorResponseResponse {
	var out ErrorResponseResponse
	out.Body = body
	out.Headers.XErrorCode = xErrorCode
	return out
}

// ErrorResponseResponse - Error response
type ErrorResponseResponse struct {
	Body    Error
	Headers struct {
		XErrorCode Maybe[int]
	}
}

func (r ErrorResponseResponse) writePostPets(w http.ResponseWriter) {
	r.Write(w, 500)
}

func (r ErrorResponseResponse) writePostShops(w http.ResponseWriter) {
	r.Write(w, 500)
}

func (r ErrorResponseResponse) Write(w http.ResponseWriter, code int) {
	if r.Headers.XErrorCode.IsSet {
		hs := []string{strconv.FormatInt(int64(r.Headers.XErrorCode.Value), 10)}
		for _, h := range hs {
			w.Header().Add("X-Error-Code", h)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeJSON(w, r.Body, "ErrorResponseResponse")
}
