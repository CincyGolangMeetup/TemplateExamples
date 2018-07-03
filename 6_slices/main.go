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

	book1 := Book{
		"978-0134190440",
		"The Go Programming Language",
		"Brian Kernighan",
		400,
	}

	book2 := Book{
		"978-0672337796",
		"Cloud Native Go",
		"Kevin Hoffman",
		256,
	}

	book3 := Book{
		"978-1491941959",
		"Introducing Go",
		"Caleb Doxsey",
		124,
	}

	books := []Book{
		book1,
		book2,
		book3,
	}

	tmp.Execute(os.Stdout, books)
}

type Book struct {
	ISBN   string
	Name   string
	Author string
	Pages  int
}
