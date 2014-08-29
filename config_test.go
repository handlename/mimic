package mimic

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	config, err := NewConfig("config.sample.json")

	if err != nil {
		t.Errorf("error: %v", err)
	}

	if len(config.Rules) != 2 {
		t.Errorf("error: unexpected length of rule.")
	}
}
