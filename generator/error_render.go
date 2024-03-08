package generator

func QueryParseError(param string) ErrorRender {
	return parseError{"query", param}
}

func HeaderParseError(param string) ErrorRender {
	return parseError{"header", param}
}

func PathParseError(param string) ErrorRender {
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
