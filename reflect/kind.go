package main

import (
	"fmt"
	"reflect"
)

func main() {
	for _, v := range []any{"hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println("string:", v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println("int:", v.Int())
		case reflect.Func:
			fmt.Println("func:", v.Kind())
		default:
			fmt.Printf("unhandled kind: %s", v.Kind())
		}
	}
}
