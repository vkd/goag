package specification

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

type SecurityRequirements []SecurityRequirement

func NewSecurityRequirements(s openapi3.SecurityRequirements, schemes SecuritySchemes) []SecurityRequirement {
	out := make([]SecurityRequirement, 0, len(s))
	for _, sr := range s {
		for k, v := range sr {
			out = append(out, NewSecurityRequirement(k, v, schemes))
			break
		}
	}
	return out
}

type SecurityRequirement struct {
	Scheme       *SecurityScheme
	Requirements []string
}

func NewSecurityRequirement(k string, v []string, schemes SecuritySchemes) SecurityRequirement {
	scheme, ok := schemes.Get("#/components/securitySchemes/" + k)
	if !ok {
		panic(fmt.Sprintf("cannot find %q security scheme", k))
	}
	return SecurityRequirement{
		Scheme:       scheme.V.Value(),
		Requirements: v,
	}
}
