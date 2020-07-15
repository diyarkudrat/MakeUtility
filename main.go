package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/bregydoc/gtranslate"
)

func main() {
	fileFlag := flag.String("file", "english-words.docx", "define input text")
	flag.Parse()

	// readFile(*fileFlag)
	translateText(*fileFlag)
}

func translateText(fileName string) {
	text := readFile(fileName)
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("en: %s | ja: %s \n", text, translated)
	// en: Hello World | ja: こんにちは世界
}

func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(fileContents))
	return string(fileContents)
}
