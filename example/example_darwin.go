//go:build darwin

package main

func osSpecific() string {
	return "hello from macOS"
}
