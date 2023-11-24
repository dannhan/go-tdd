package main

import "fmt"

const (
	spanish    = "Spanish"
	french     = "French"
	indonesian = "Indonesian"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	indonesianHelloPrefix = "Halo, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case indonesian:
		prefix = indonesianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
