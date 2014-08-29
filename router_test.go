package mimic

import (
	"testing"
)

func TestRouting(t *testing.T) {
	c, _ := NewConfig("config.sample.json")
	r := routing(c.Rules)

	if r == nil {
		t.Error("router is nil.")
	}

	for _, rule := range c.Rules {
		route := r.Get(rule.String())

		if route == nil {
			t.Errorf("route not registered: %s", route)
		}
	}
}
