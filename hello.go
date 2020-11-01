package main

import "fmt"

const (
	french                = "French"
	spanish               = "Spanish"
	portuguese            = "Portuguese"
	portugueseHelloPrefix = "Bom dia, "
	frenchHelloPrefix     = "Bonjour, "
	spanishHelloPrefix    = "Hola, "
	englishHelloPrefix    = "Hello, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Itanor", "Portuguese"))
}
