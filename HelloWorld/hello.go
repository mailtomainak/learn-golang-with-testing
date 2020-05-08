package main

import "fmt"

const englishHelloPrefix = "Hello "
const frenchHelloPrefix = "Bonjour "
const spanishHelloPrefix = "Hola "

func setPrefix(language string) (prefix string) {
	switch language {

	case "ES":
		prefix = spanishHelloPrefix

	case "FR":
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix

	}
	return
}

// Hello Returns the hello string
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return setPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Chris", "ES"))
}
