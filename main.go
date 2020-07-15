package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/bregydoc/gtranslate"
)

type pageData struct {
	Content string
	Title   string
}

func main() {
	fileFlag := flag.String("file", "english-words.docx", "define input text")
	flag.Parse()

	runFile(*fileFlag)
}

func runFile(fileFlag string) {

	var fileName string = fileFlag
	var fileTypes = [4]string{".txt", ".docx", ".odt", ".rtf"}

	for _, ext := range fileTypes {
		if fileName[strings.Index(fileFlag, "."):len(fileFlag)] != ext {
			return
		}
	}

	fileName = fileName[0:strings.Index(fileFlag, ".")] + ".txt"
	var data string = readFile(fileFlag)
	renderText("template.tmpl", data, fileName)

}

func translateText(txtData string) {

	translated, err := gtranslate.TranslateWithParams(
		txtData,
		gtranslate.TranslationParams{
			From: "en",
			To:   "ja",
		},
	)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("en: %s | ja: %s \n", txtData, translated)
	// en: Hello World | ja: こんにちは世界
	return string(translated)
}

func renderText(tPath, textData, fileName string) {
	paths := []string{
		tPath,
	}

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	text, err := template.New(tPath).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}

	txtTranslated, err := translateText(textData)

	originName := fileName[0:strings.Index(fileName, ".")]

	err = text.Execute(file, pageData{txtTranslated, originName})
	if err != nil {
		panic(err)
	}

	file.Close()
}

func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(fileContents))
	return string(fileContents)
}
