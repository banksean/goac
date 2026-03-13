package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Printf("Operating System: %s\n", os)

	switch os {
	case "windows":
		fmt.Println("Running on Windows")
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		fmt.Println("Running on Linux")
	default:
		fmt.Printf("Running on %s\n", os)
	}

	fmt.Printf("result of osSpecific() call: %s\n", osSpecific())
}
