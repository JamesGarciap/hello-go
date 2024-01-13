package main

import "fmt"

func Hello(name string) string {
	if name != "" {
		return "Hello, " + name
	}

	return "Hello, World!"
}

// main is our application entry point
// private methods names start with lowercase, and so public ones start with uppercase
func main() {
	fmt.Println("Holis")
}
