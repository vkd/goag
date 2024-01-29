package specification

import "github.com/getkin/kin-openapi/openapi3"

type Info struct {
	Title       string
	Description string
	Version     string
}

func NewInfo(s *openapi3.Info) Info {
	if s == nil {
		return Info{}
	}
	return Info{
		Title:       s.Title,
		Description: s.Description,
		Version:     s.Version,
	}
}
