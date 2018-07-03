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

	colors := map[string]string{
		"white": "#FFFFFF",
		"black": "#000000",
		"red":   "#FF0000",
		"blue":  "#0000FF",
	}

	tmp.Execute(os.Stdout, colors)
}
