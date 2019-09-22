package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const englishHelloPrefix = "Hello, "
const thaiHelloPrefix = "สวัสดี, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

const thai = "Thai"
const spanish = "Spanish"
const french = "French"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case thai:
		prefix = thaiHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
