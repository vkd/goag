package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type Security struct {
	Requirement []string
	Scheme      openapi3.SecurityScheme
}

func GetSecurity(secSchemes openapi3.SecuritySchemes, secRequirements openapi3.SecurityRequirements) ([][]Security, error) {
	var out [][]Security

	for _, sr := range secRequirements {
		var sOut []Security
		for scheme, reqs := range sr {
			ref, ok := secSchemes[scheme]
			if !ok {
				return nil, fmt.Errorf("security requirement key %q: not found in 'components.securitySchemes'", scheme)
			}
			if ref.Value == nil {
				return nil, fmt.Errorf("security scheme %q is empty", scheme)
			}
			sOut = append(sOut, Security{
				Requirement: reqs,
				Scheme:      *ref.Value,
			})
		}
		out = append(out, sOut)
	}

	return out, nil
}
