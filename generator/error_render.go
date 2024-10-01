package generator

type parseParamError struct {
	InType    string
	Parameter string
}

func (p parseParamError) Wrap(reason string, errVar string) string {
	return `ErrParseParam{In: "` + p.InType + `", Parameter: "` + p.Parameter + `", Reason: "` + reason + `", Err: ` + errVar + `}`
}

func (p parseParamError) New(reason string) string {
	return `ErrParseParam{In: "` + p.InType + `", Parameter: "` + p.Parameter + `", Reason: "` + reason + `"}`
}

type newError struct{}

func (p newError) Wrap(reason string, errVar string) string {
	return `fmt.Errorf("` + reason + `: %w", ` + errVar + `)`
}

func (p newError) New(reason string) string {
	return `errors.New("` + reason + `")`
}

type prefixError struct {
	prefix string
}

func (p prefixError) Wrap(reason string, errVar string) string {
	return `fmt.Errorf("` + p.prefix + ": " + reason + `: %w", ` + errVar + `)`
}

func (p prefixError) New(reason string) string {
	return `errors.New("` + p.prefix + ": " + reason + `")`
}

type wrappedError struct {
	Orig  ErrorRender
	Inner ErrorRender
}

func (p wrappedError) Wrap(reason string, errVar string) string {
	return p.Orig.Wrap(reason, p.Inner.Wrap(reason, errVar))
}

func (p wrappedError) New(reason string) string {
	return p.Orig.Wrap(reason, p.Inner.New(reason))
}

type returnsArgs struct {
	returns string
	inner   ErrorRender
}

func (e returnsArgs) Wrap(reason string, errVar string) string {
	var inner string
	if e.inner != nil {
		inner = e.inner.Wrap(reason, errVar)
	} else {
		inner = errVar
	}
	if e.returns != "" {
		return e.returns + ", " + inner
	}
	return inner
}

func (e returnsArgs) New(reason string) string {
	var inner string
	if e.inner != nil {
		inner = e.inner.New(reason)
	} else {
		inner = reason
	}
	if e.returns != "" {
		return e.returns + ", " + inner
	}
	return inner
}
