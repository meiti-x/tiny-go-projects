package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[string]string{
	"en": "Hello World",
	"fr": "Bonjour le monde",
}

func main() {
	lan := flag.String("lang", "", "the required language, en")

	flag.Parse()

	message := greeting(*lan)
	fmt.Println(message)
}
func greeting(lang string) string {
	greet, ok := phrasebook[lang]

	if !ok {
		return fmt.Sprintf("Unknown language: %q", lang)
	}
	return greet
}
