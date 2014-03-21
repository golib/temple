package html

import (
	"testing"
)

func assetBool(got, expected bool, t *testing.T) {
	if got != expected {
		t.Fatalf("Expected {%s} but got {%s}.", expected, got)
	}
}

func assetInt(got, expected int, t *testing.T) {
	if got != expected {
		t.Fatalf("Expected {%s} but got {%s}.", expected, got)
	}
}

func assetString(got, expected string, t *testing.T) {
	if got != expected {
		t.Fatalf("Expected {%s} but got {%s}.", expected, got)
	}
}
