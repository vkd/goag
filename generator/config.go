package generator

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Maybe struct {
		Import string `yaml:"import"`
		Type   string `yaml:"type"`
	} `yaml:"maybe"`
	Nullable struct {
		Import string `yaml:"import"`
		Type   string `yaml:"type"`
	} `yaml:"nullable"`

	Cors struct {
		Enable bool `yaml:"enable"`
	} `yaml:"cors"`

	CustomTypes struct {
		Ignore bool `yaml:"ignore"`
	} `yaml:"customTypes"`

	CustomImports []struct {
		Value string `yaml:"value"`
		Alias string `yaml:"alias"`
	} `yaml:"imports"`
}

func LoadConfig(filepath string) (zero Config, _ error) {
	bs, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return zero, nil
		}
		return zero, fmt.Errorf("read cfg file: %w", err)
	}

	var c Config
	err = yaml.Unmarshal(bs, &c)
	if err != nil {
		return zero, fmt.Errorf("unmarshal yaml config: %w", err)
	}

	return c, nil
}

func (c Config) Imports() Imports {
	out := make(Imports, 0, len(c.CustomImports))
	for _, ci := range c.CustomImports {
		out = append(out, NewImport(ci.Value, ci.Alias))
	}
	return out
}
