package generator

type Render interface {
	Render() (string, error)
}

type RenderFunc func() (string, error)

func (r RenderFunc) Render() (string, error) { return r() }

type StringRender string

func (s StringRender) Render() (string, error) { return string(s), nil }

type QuotedRender string

func (s QuotedRender) Render() (string, error) { return string(`"` + s + `"`), nil }

type ErrorRender interface {
	Wrap(message string) string
	New(message string) string
}

// Parser parses 'string' to '<type>'.
type Parser interface {
	RenderParser(from, to Render, mkErr ErrorRender) (string, error)
}

type ParserFunc func(from, to Render, mkErr ErrorRender) (string, error)

func (p ParserFunc) RenderParser(from, to Render, mkErr ErrorRender) (string, error) {
	return p(from, to, mkErr)
}

// Formatter formats 'string' from '<type>'.
type Formatter interface {
	RenderFormat(from Render) (string, error)
}

type FormatterFunc func(from Render) (string, error)

func (f FormatterFunc) RenderFormat(from Render) (string, error) { return f(from) }

type Renders []Render

func (r Renders) Render() (string, error) {
	return ExecuteTemplate("Renders", r)
}
