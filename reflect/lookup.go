package main

import (
	"fmt"
	"reflect"
)

func main() {
	type S struct {
		F0 string `alias:"field_0"`
		F1 string `alias:""`
		F2 string
		F3 int
		F4 struct{}
	}

	s := S{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		typ := field.Type.Kind()
		fmt.Printf("%s %s ", field.Name, typ)
		fmt.Print("[alias ")
		if alias, ok := field.Tag.Lookup("alias"); ok {
			if alias == "" {
				fmt.Print("(blank)")
			} else {
				fmt.Print(alias)
			}
		} else {
			fmt.Print("(not specified)")
		}
		fmt.Println("]")
	}
}
