package specification

import "github.com/getkin/kin-openapi/openapi3"

type Server struct {
	URL         string
	Description string
	Variables   Map[ServerVariable]
}

func NewServer(s *openapi3.Server) Server {
	return Server{
		URL:         s.URL,
		Description: s.Description,
		Variables:   NewMap(s.Variables, NewServerVariable),
	}
}

func NewServers(ss openapi3.Servers) []Server {
	out := make([]Server, 0, len(ss))
	for _, s := range ss {
		out = append(out, NewServer(s))
	}
	return out
}

type ServerVariable struct {
	Enum        []string
	Default     string
	Description string
}

func NewServerVariable(sv *openapi3.ServerVariable) ServerVariable {
	enums := make([]string, 0, len(sv.Enum))
	for _, e := range sv.Enum {
		enums = append(enums, e.(string))
	}
	return ServerVariable{
		Enum:        enums,
		Default:     sv.Default.(string),
		Description: sv.Description,
	}
}
