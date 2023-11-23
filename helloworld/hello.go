package main

import (
	"fmt"
)

// group all constants into single declaration
const (
	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	italianHelloPrefix    = "Buongiorno, "
	portugueseHelloPrefix = "Olá, "
)

// public function: starts with a capital letter
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// private function: starts with a lowercase letter
func greetingPrefix(language string) (prefix string) {

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "Italian":
		prefix = italianHelloPrefix
	case "Portuguese":
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
	// return the named return value "prefix"
	// "prefix" initialized as empty and assigned inside func
}

func main() {
	fmt.Println(Hello("Amélie", "French"))
	fmt.Println(Hello("Michelangelo", "Italian"))
	fmt.Println(Hello("Elizabeth", "English"))
}
