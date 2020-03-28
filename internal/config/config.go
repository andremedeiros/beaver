package config

import "gopkg.in/yaml.v3"

type Config struct {
	Name     string
	Language string
	Tasks    map[string]Task `yaml:",omitempty"`
}

func Decode(b []byte) (*Config, error) {
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Config) Encode() ([]byte, error) {
	return yaml.Marshal(c)
}

func (c *Config) Validate() error {
	return nil
}
