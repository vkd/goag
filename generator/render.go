package generator

import (
	"bytes"
	"fmt"
	"text/template"
)

type Render interface {
	String() (string, error)
}

func String(tm *template.Template, data interface{}) (string, error) {
	var bs bytes.Buffer
	err := tm.Execute(&bs, data)
	if err != nil {
		return "", fmt.Errorf("to string: %w", err)
	}
	return bs.String(), nil
}

func Bytes(r Render) ([]byte, error) {
	out, err := r.String()
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return []byte(out), nil
}
