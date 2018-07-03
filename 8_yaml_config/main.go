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
	tmp = template.Must(template.ParseFiles("tpl.yaml"))
}

func main() {

	// You could image pulling these values from a database
	// or passing them as command line flags
	c := Config{
		"nginx",
		"10.40.123.65",
		"nginx:latest",
		8080,
	}
	// This time we are creating a file and write there instead.
	f, err := os.Create("out.yaml")
	if err != nil {
		log.Printf("File could not be created: %v", err)
	}
	defer f.Close()

	// template.Execute takes the io.Writer interface
	err = tmp.Execute(f, c)
	if err != nil {
		log.Fatal(err)
	}
}

// Config is an object that contains the values passed to the template.
type Config struct {
	Name      string
	ClusterIP string
	Image     string
	Port      int
}
