package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type SecurityRequirements []SecurityRequirement

func NewSecurityRequirements(s openapi3.SecurityRequirements, schemes SecuritySchemes) ([]SecurityRequirement, error) {
	out := make([]SecurityRequirement, 0, len(s))
	for _, sr := range s {
		for k, v := range sr {
			ss, err := NewSecurityRequirement(k, v, schemes)
			if err != nil {
				return nil, fmt.Errorf("new security requirements %q: %w", k, err)
			}
			out = append(out, ss)
			break
		}
	}
	return out, nil
}

type SecurityRequirement struct {
	Scheme       *SecurityScheme
	Requirements []string
}

func NewSecurityRequirement(k string, v []string, schemes SecuritySchemes) (zero SecurityRequirement, _ error) {
	scheme, ok := schemes.Get("#/components/securitySchemes/" + k)
	if !ok {
		return zero, fmt.Errorf("cannot find %q security scheme", k)
	}
	return SecurityRequirement{
		Scheme:       scheme.V.Value(),
		Requirements: v,
	}, nil
}
