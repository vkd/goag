package generator

type HandlerOld struct {
	// client

	// deprecated
	Name        OperationName
	Description string
	Summary     string

	BasePathPrefix string

	CanParseError bool

	ResponserInterfaceName string

	IsJWT bool

	Parameters struct {
		Query   []Param
		Path    []Param
		Headers []Param

		PathParsers []Templater
	}

	Body struct {
		TypeName Templater
	}

	IsWriteJSONFunc bool

	Responses []Templater
}

func (h HandlerOld) Render() (string, error) {
	return h.Execute()
}

func (h HandlerOld) HandlerFuncName() string { return string(h.Name) + "HandlerFunc" }

func (h HandlerOld) Execute() (string, error) { return templates.ExecuteTemplate("Handler", h) }

type Param struct {
	Field  Templater
	Parser Templater
}
