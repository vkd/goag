package generator

type HandlersFile struct {
	Handlers []HandlerOld

	IsWriteJSONFunc bool
}

func (g *Generator) HandlersFile(hs []HandlerOld, isJSON bool) (Templater, error) {
	file := HandlersFile{
		Handlers:        hs,
		IsWriteJSONFunc: isJSON,
	}

	return g.goFile([]Import{
		"encoding/json",
		"fmt",
		"io",
		"log",
		"net/http",
		"strconv",
		"strings",
	}, RenderFunc(OldTemplater(file).String)), nil
}

func (h HandlersFile) Execute() (string, error) { return templates.ExecuteTemplate("HandlersFile", h) }

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

func (h HandlerOld) HandlerFuncName() string { return string(h.Name) + "HandlerFunc" }

func (h HandlerOld) Execute() (string, error) { return templates.ExecuteTemplate("Handler", h) }

type Param struct {
	Field  Templater
	Parser Templater
}
