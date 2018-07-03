package main

import (
	"log"
	"os"

	"github.com/alecthomas/template"
)

var tmp *template.Template

// This makes sure that the template only gets processed once
// template.Must does the error handling for us.
func init() {
	tmp = template.Must(template.ParseFiles("tpl.txt"))
}

func main() {
	// This time we are creating a file and write there instead.
	f, err := os.Create("out.txt")
	if err != nil {
		log.Printf("File could not be created: %v", err)
	}
	defer f.Close()

	// template.Execute takes the io.Writer interface
	err = tmp.Execute(f, nil)
	if err != nil {
		log.Fatal(err)
	}
}
