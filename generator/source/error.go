package source

type ErrorWrapper interface {
	Wrap(reason string) string
}

type ErrorBuilder interface {
	New(reason string) string
}

type ParseError struct {
	In, Parameter string
}

func (p ParseError) Wrap(reason string) string {
	return `ErrParseParam{In: "` + p.In + `", Parameter: "` + p.Parameter + `", Reason: "` + reason + `", Err: err}`
}

func (p ParseError) New(reason string) string {
	return `ErrParseParam{In: "` + p.In + `", Parameter: "` + p.Parameter + `", Reason: "` + reason + `"}`
}
