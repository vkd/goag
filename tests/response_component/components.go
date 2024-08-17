package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ------------------------
//         Schemas
// ------------------------

type Error struct {
	Message string `json:"message"`
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
	write([]byte(`"message":`))
	writeJSON(c.Message)

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

// ------------------------------
//         Responses
// ------------------------------

func NewErrorResponse(code int, body Error) ErrorResponse {
	var out ErrorResponse
	out.Code = code
	out.Body = body
	return out
}

// ErrorResponse - Error output response
type ErrorResponse struct {
	Code int
	Body Error
}

func (r ErrorResponse) writeGetV2Pet(w http.ResponseWriter) {
	r.Write(w)
}

func (r ErrorResponse) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	writeJSON(w, r.Body, "ErrorResponse")
}

func NewPetResponse(body Pet) PetResponse {
	var out PetResponse
	out.Body = body
	return out
}

// PetResponse - Pet output response
type PetResponse struct {
	Body Pet
}

func (r PetResponse) writeGetPet(w http.ResponseWriter) {
	r.Write(w, 200)
}

func (r PetResponse) writeGetV2Pet(w http.ResponseWriter) {
	r.Write(w, 201)
}

func (r PetResponse) writeGetV3Pet(w http.ResponseWriter) {
	r.Write(w, 202)
}

func (r PetResponse) Write(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	writeJSON(w, r.Body, "PetResponse")
}

// NewPet2Response - Pet output response
func NewPet2Response(body Pet) Pet2Response {
	return NewPetResponse(body)
}

// Pet2Response - Pet output response
type Pet2Response = PetResponse

// NewPet3Response - Pet output response
func NewPet3Response(body Pet) Pet3Response {
	return NewPet2Response(body)
}

// Pet3Response - Pet output response
type Pet3Response = Pet2Response
