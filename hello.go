package main

import "fmt"

const spanish = "Spanish"
const swedish = "Swedish"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const swedishPrefix = "Hej, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix

	switch language {
	case spanish:
		prefix = spanishPrefix
	case swedish:
		prefix = swedishPrefix
	}

	return prefix + name
}

// main is our application entry point
// private methods names start with lowercase, and so public ones start with uppercase
func main() {
	fmt.Println("Holis")
}
