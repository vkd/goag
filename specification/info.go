package specification

import "github.com/getkin/kin-openapi/openapi3"

type Info struct {
	Title          string
	Summary        string
	Description    string
	TermsOfService string
	Contact        Contact
	License        License
	Version        string
}

func NewInfo(s *openapi3.Info) Info {
	if s == nil {
		return Info{}
	}
	return Info{
		Title: s.Title,
		// Summary: s.ExtensionProps.Extensions["summary"],
		Description:    s.Description,
		TermsOfService: s.TermsOfService,
		Contact:        NewContact(s.Contact),
		License:        NewLicense(s.License),
		Version:        s.Version,
	}
}

type Contact struct {
	Name  string
	URL   string
	Email string
}

func NewContact(c *openapi3.Contact) Contact {
	if c == nil {
		return Contact{}
	}
	return Contact{
		Name:  c.Name,
		URL:   c.URL,
		Email: c.Email,
	}
}

type License struct {
	Name       string
	Identifier string
	URL        string
}

func NewLicense(l *openapi3.License) License {
	if l == nil {
		return License{}
	}
	return License{
		Name: l.Name,
		// Identifier: l.ExtensionProps.Extensions["identifier"],
		URL: l.URL,
	}
}
