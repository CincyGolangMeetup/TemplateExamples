package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/template"
)

func main() {
	// NOTE: You can pass any number of files to ParseFiles
	tmp, err := template.ParseFiles("tpl.txt", "cincinnati.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}

	tmp.Execute(os.Stdout, nil)
}
