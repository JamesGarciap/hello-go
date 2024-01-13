package main

import "fmt"

const spanish = "Spanish"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix

	if language == spanish {
		prefix = spanishPrefix
	}

	return prefix + name
}

// main is our application entry point
// private methods names start with lowercase, and so public ones start with uppercase
func main() {
	fmt.Println("Holis")
}
