package main

import "fmt"

func PluginPrint(s string) {
	fmt.Printf("PluginPrint: %s\n", s)
}

func main() {
	fmt.Println("main from plugin.go")
}
