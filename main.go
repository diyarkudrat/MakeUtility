package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

type pageData struct {
	Content string
	Title   string
}

func main() {
	fileFlag := flag.String("file", "english-words.docx", "define input text")
	flag.Parse()

	// fmt.Println(fileFlag)
	runFile(*fileFlag, "txt_dir/")
}

func containsExt(fileFlag string) bool {

	var fileName string = fileFlag
	var fileTypes = [4]string{".docx", ".txt", ".odt", ".rtf"}
	extension := fileName[strings.Index(fileFlag, "."):len(fileFlag)]

	for _, ext := range fileTypes {
		if extension != ext {
			return true
		}
	}
	return false
}

func runFile(fileFlag, directory string) {
	var fileName string = fileFlag

	if !containsExt(fileFlag) {
		fmt.Println("Provide valid file")
	}

	fileName = fileName[0:strings.Index(fileFlag, ".")] + ".html"
	var data string = readFile(directory + fileFlag)
	fmt.Println(data)
	renderText("template.tmpl", data, fileName)

}

func translateText(txtData string) string {

	translated, err := gtranslate.Translate(txtData, language.English, language.Spanish)
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

	file, err := os.Create("templates/" + fileName)
	if err != nil {
		panic(err)
	}

	text, err := template.New(tPath).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}

	txtTranslated := translateText(textData)
	fmt.Println(txtTranslated)

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
