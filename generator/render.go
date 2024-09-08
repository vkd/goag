package generator

type Render interface {
	Render() (string, error)
}

type RenderFunc func() (string, error)

func (r RenderFunc) Render() (string, error) { return r() }

type GoTypeRender interface {
	RenderGoType() (string, error)
}

type GoTypeRenderFunc func() (string, error)

func (r GoTypeRenderFunc) Render() (string, error) { return r() }

type StringRender string

func (s StringRender) Render() (string, error) { return string(s), nil }

type ErrorRender interface {
	Wrap(reason string, errVar string) string
	New(message string) string
}

// Parser parses 'string' to '<type>'.
type Parser interface {
	ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error)
	IsMultivalue() bool
	ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error)
}

type SingleValue struct{}

func (SingleValue) IsMultivalue() bool { return false }

type Multivalue struct{}

func (Multivalue) IsMultivalue() bool { return true }

type ParserFunc func(to, from string, isNew bool, mkErr ErrorRender) (string, error)

func (p ParserFunc) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return p(to, from, isNew, mkErr)
}

// Formatter formats 'string' from '<type>'.
type Formatter interface {
	RenderFormat(from string) (string, error)
	RenderFormatStrings(to, from string, isNew bool) (string, error)
}

type FormatterFunc func(from string) (string, error)

func (f FormatterFunc) RenderFormat(from string) (string, error) { return f(from) }

type Renders []Render

func (r Renders) Render() (string, error) {
	return ExecuteTemplate("Renders", r)
}
