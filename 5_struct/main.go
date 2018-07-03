package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/template"
)

func main() {
	// NOTE: You can pass any number of files to ParseFiles
	tmp, err := template.ParseFiles("tpl.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}

	book := Book{
		"978-0134190440",
		"The Go Programming Language",
		"Brian Kernighan",
		400,
	}

	tmp.Execute(os.Stdout, book)
}

type Book struct {
	ISBN   string
	Name   string
	Author string
	Pages  int
}
