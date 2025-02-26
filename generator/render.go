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
	ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error)
}

type NotImplementedParser struct{}

func (NotImplementedParser) ParseString(_, _ string, _ bool, _ ErrorRender) (string, error) {
	panic("ParseString: not implemented")
}

func (NotImplementedParser) ParseStrings(_, _ string, _ bool, _ ErrorRender) (string, error) {
	panic("ParseStrings: not implemented")
}

type ParserFunc func(to, from string, isNew bool, mkErr ErrorRender) (string, error)

func (p ParserFunc) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return p(to, from, isNew, mkErr)
}

// Formatter formats 'string' from '<type>'.
type Formatter interface {
	RenderFormat(from string) (string, error)
	RenderFormatStrings(to, from string, isNew bool) (string, error)
}

type NotImplementedFormatter struct{}

func (NotImplementedFormatter) RenderFormat(_ string) (string, error) {
	panic("RenderFormat: not implemented")
}

func (NotImplementedFormatter) RenderFormatStrings(_, _ string, _ bool) (string, error) {
	panic("RenderFormatStrings: not implemented")
}

type FormatterFunc func(from string) (string, error)

func (f FormatterFunc) RenderFormat(from string) (string, error) { return f(from) }

type Renders []Render

func (r Renders) Render() (string, error) {
	return ExecuteTemplate("Renders", r)
}
