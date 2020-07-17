package main

import (
	"flag"
	"fmt"
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
	langFlag := flag.String("language", "", "define what language text is in")
	translatelLangFlag := flag.String("translated", "", "define what language to translate to")
	flag.Parse()

	// fmt.Println(fileFlag)
	runFile(*fileFlag, *langFlag, *translatelLangFlag, "txt_dir/")
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

func runFile(fileFlag, langFlag, translatelLangFlag, directory string) {
	var fileName string = fileFlag

	if !containsExt(fileFlag) {
		fmt.Println("Provide valid file")
	}

	fileName = fileName[0:strings.Index(fileFlag, ".")] + ".html"
	var data string = readFile(directory + fileFlag)
	fmt.Println(data)
	renderText("template.tmpl", data, fileName, langFlag, translatelLangFlag)

}

func translateText(txtData, langFlag, translatelLangFlag string) string {

	// fmt.Printf("Name: %v\nPrice: %v\n\n", item.Name, item.Price)
	translated, err := gtranslate.TranslateWithParams(
		txtData,
		gtranslate.TranslationParams{
			From: langFlag,
			To:   translatelLangFlag,
		},
	)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("en: %s | ja: %s \n", txtData, translated)
	// en: Hello World | ja: こんにちは世界
	return string(translated)
}

func renderText(tPath, textData, fileName, langFlag, translatelLangFlag string) {
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

	txtTranslated := translateText(textData, langFlag, translatelLangFlag)
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
