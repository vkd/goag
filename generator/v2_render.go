package generator

type Render interface {
	Render() (string, error)
}

type RenderFunc func() (string, error)

func (r RenderFunc) Render() (string, error) { return r() }

type StringRender string

func (s StringRender) Render() (string, error) { return string(s), nil }

type ErrorRender interface {
	Wrap(message string) string
	New(message string) string
}

type Parser interface {
	RenderParser(from, to Render, mkErr ErrorRender) (string, error)
}

type ParserFunc func(from, to Render, mkErr ErrorRender) (string, error)

func (p ParserFunc) RenderParser(from, to Render, mkErr ErrorRender) (string, error) {
	return p(from, to, mkErr)
}

type Formatter interface {
	RenderFormat(from Render) (string, error)
}

type FormatterFunc func(from Render) (string, error)

func (f FormatterFunc) RenderFormat(from Render) (string, error) { return f(from) }
