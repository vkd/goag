package test

import (
	"encoding/json"
	"fmt"
)

// ------------------------
//         Schemas
// ------------------------

type Pet struct {
	Custom               PetCustom                  `json:"custom"`
	Name                 string                     `json:"name"`
	AdditionalProperties map[string]json.RawMessage `json:"-"`
}

var _ json.Marshaler = (*Pet)(nil)

func (c *Pet) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for k, v := range c.AdditionalProperties {
		m[k] = v
	}
	m["custom"] = c.Custom
	m["name"] = c.Name
	return json.Marshal(m)
}

var _ json.Unmarshaler = (*Pet)(nil)

func (c *Pet) UnmarshalJSON(bs []byte) error {
	m := make(map[string]json.RawMessage)
	err := json.Unmarshal(bs, &m)
	if err != nil {
		return fmt.Errorf("raw key/value map: %w", err)
	}
	if v, ok := m["custom"]; ok {
		err = json.Unmarshal(v, &c.Custom)
		if err != nil {
			return fmt.Errorf("'custom' field: %w", err)
		}
		delete(m, "custom")
	}
	if v, ok := m["name"]; ok {
		err = json.Unmarshal(v, &c.Name)
		if err != nil {
			return fmt.Errorf("'name' field: %w", err)
		}
		delete(m, "name")
	}
	for k, bs := range m {
		var v json.RawMessage
		err = json.Unmarshal(bs, &v)
		if err != nil {
			return fmt.Errorf("additional property %q: %w", k, err)
		}
		c.AdditionalProperties[k] = v
	}
	return nil
}

type PetCustom = json.RawMessage

type Pets []Pet
