package generator

import (
	"fmt"
	"strings"
)

type BoolType struct{}

var _ Parser = BoolType{}

func (b BoolType) RenderParser(from, to Render, mkErr ErrorRender) (string, error) {
	return ExecuteTemplate("BoolParser", struct {
		From  Render
		To    Render
		MkErr ErrorRender
	}{from, to, mkErr})
}

var _ Formatter = BoolType{}

func (b BoolType) RenderFormat(from Render) (string, error) {
	return ExecuteTemplate("BoolFormat", struct {
		From Render
	}{from})
}

type IntType struct {
	BitSize int
}

var _ Parser = IntType{}

func (i IntType) RenderParser(from, to Render, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 0:
		return ExecuteTemplate("IntParser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	case 32:
		return ExecuteTemplate("Int32Parser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	case 64:
		return ExecuteTemplate("Int64Parser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	default:
		return "", fmt.Errorf("unsupported int bit size %d", i.BitSize)
	}
}

var _ Formatter = IntType{}

func (i IntType) RenderFormat(from Render) (string, error) {
	switch i.BitSize {
	case 0:
		return ExecuteTemplate("IntFormat", struct {
			From Render
		}{from})
	case 32:
		return ExecuteTemplate("Int32Format", struct {
			From Render
		}{from})
	case 64:
		return ExecuteTemplate("Int64Format", struct {
			From Render
		}{from})
	default:
		return "", fmt.Errorf("unsupported int bit size %d", i.BitSize)
	}
}

type FloatType struct {
	BitSize int
}

var _ Parser = FloatType{}

func (i FloatType) RenderParser(from, to Render, mkErr ErrorRender) (string, error) {
	switch i.BitSize {
	case 32:
		return ExecuteTemplate("Float32Parser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	case 64:
		return ExecuteTemplate("Float64Parser", struct {
			From  Render
			To    Render
			MkErr ErrorRender
		}{from, to, mkErr})
	default:
		return "", fmt.Errorf("unsupported float bit size %d", i.BitSize)
	}
}

var _ Formatter = FloatType{}

func (i FloatType) RenderFormat(from Render) (string, error) {
	switch i.BitSize {
	case 32:
		return ExecuteTemplate("Float32Format", struct {
			From Render
		}{from})
	case 64:
		return ExecuteTemplate("Float64Format", struct {
			From Render
		}{from})
	default:
		return "", fmt.Errorf("unsupported float bit size %d", i.BitSize)
	}
}

type StringType struct{}

func (_ StringType) RenderFormat(from Render) (string, error) {
	return from.Render()
}

type CustomType string

func NewCustomType(s string) CustomType {
	var customImport, customType string = "", s
	slIdx := strings.LastIndex(s, "/")
	if slIdx >= 0 {
		customImport = s[:slIdx]
		customType = s[slIdx+1:]

		dotIdx := strings.LastIndex(s, ".")
		if dotIdx >= 0 {
			customImport = s[:dotIdx]
		}
	}

	CustomImports = append(CustomImports, customImport)
	return CustomType(customType)
}

func (c CustomType) RenderFormat(from Render) (string, error) {
	return ExecuteTemplate("CustomTypeFormat", struct {
		From Render
	}{From: from})
}
