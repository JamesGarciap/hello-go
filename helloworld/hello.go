package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	swedish = "Swedish"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	swedishPrefix = "Hej, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case swedish:
		prefix = swedishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

// main is our application entry point
// private methods names start with lowercase, and so public ones start with uppercase
func main() {
	fmt.Println("Holis")
}
