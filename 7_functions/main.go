package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/template"
)

func main() {

	funcMap := template.FuncMap{
		"title":        strings.Title,
		"upper":        strings.ToUpper,
		"removeLetter": removeLetter,
	}

	tmp, err := template.New("tpl.txt").Funcs(funcMap).ParseFiles("tpl.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}

	name := "cincinnati go meetup"

	tmp.Execute(os.Stdout, name)
}

func removeLetter(s, l string) string {
	n := strings.Replace(s, l, "", -1)
	return n
}
