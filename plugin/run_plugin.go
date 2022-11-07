package main

import (
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("./plugin")
	if err != nil {
		log.Fatal(err)
	}
	symb, err := p.Lookup("PluginPrint")
	if err != nil {
		log.Fatal(err)
	}
	f, ok := symb.(func(string))
	if !ok {
		log.Fatal("symb is not func(string)")
	}
	f("call from run_plugin.go!")
}
