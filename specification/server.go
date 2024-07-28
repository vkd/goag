package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Server struct {
	URL         string
	Description string
	Variables   Map[ServerVariable]
}

func NewServer(s *openapi3.Server) (zero Server, _ error) {
	variables, err := NewMap[ServerVariable, *openapi3.ServerVariable](s.Variables, func(sv *openapi3.ServerVariable) (ServerVariable, error) {
		return NewServerVariable(sv), nil
	})
	if err != nil {
		return zero, fmt.Errorf("new variables: %w", err)
	}
	return Server{
		URL:         s.URL,
		Description: s.Description,
		Variables:   variables,
	}, nil
}

func NewServers(ss openapi3.Servers) ([]Server, error) {
	out := make([]Server, 0, len(ss))
	for _, s := range ss {
		server, err := NewServer(s)
		if err != nil {
			return nil, fmt.Errorf("new server: %w", err)
		}
		out = append(out, server)
	}
	return out, nil
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
