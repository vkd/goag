package source

type ErrorWrapper = ErrorBuilder

type ErrorBuilder interface {
	Wrap(reason string) string
	New(reason string) string
}

func QueryParseError(param string) ErrorBuilder {
	return parseError{"query", param}
}

func HeaderParseError(param string) ErrorBuilder {
	return parseError{"header", param}
}

func PathParseError(param string) ErrorBuilder {
	return parseError{"path", param}
}

type parseError struct {
	InType    string
	Parameter string
}

func (p parseError) Wrap(reason string) string {
	return `ErrParseParam{In: "` + p.InType + `", Parameter: "` + p.Parameter + `", Reason: "` + reason + `", Err: err}`
}

func (p parseError) New(reason string) string {
	return `ErrParseParam{In: "` + p.InType + `", Parameter: "` + p.Parameter + `", Reason: "` + reason + `"}`
}
