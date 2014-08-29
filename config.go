package mimic

import (
	"encoding/json"
	"io/ioutil"
)

// Config stores configuration loaded from config.json
type Config struct {
	Port  int    `json:"port"`
	Rules []Rule `json:"rules"`
}

// Rule stores configuration of a routing rule.
type Rule struct {
	Method  string
	Path    string
	Content string
}

// NewConfig returns Config stores cnofigration
// read from json at configPath.
func NewConfig(configPath string) (Config, error) {
	content, err := ioutil.ReadFile(configPath)

	if err != nil {
		return Config{}, err
	}

	c := Config{}
	err = json.Unmarshal([]byte(content), &c)

	if err != nil {
		return Config{}, err
	}

	if c.Port == 0 {
		c.Port = 3390
	}

	return c, nil
}
