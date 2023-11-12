package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const helloSuffix = "!"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name + helloSuffix
	}

	return englishHelloPrefix + name + helloSuffix
}

func main() {
	fmt.Println(Hello("world", ""))
}
