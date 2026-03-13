//go:build linux

package main

import "testing"

func Test_osSpecific(t *testing.T) {
	expected := "hello from linux"
	if result := osSpecific(); result != expected {
		t.Errorf("result: %q, got: %q", result, expected)
	}
}
