package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	const text = "<<.Greeting>> {{.Name}}"

	data := struct {
		Greeting string
		Name     string
	}{
		Greeting: "Hello",
		Name:     "Joe",
	}

	t := template.Must(template.New("tpl").Delims("<<", ">>").Parse(text))

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
