package client

import (
	"testing"
)

func TestF1(t *testing.T) {
	err := F1()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
