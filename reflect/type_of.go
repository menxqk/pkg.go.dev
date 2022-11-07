package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// As interface types are only used for static typing, a
	// common idiom to find the reflection Type for an interface
	// type Foo is to use a *Foo value..
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Printf("%v\n", writerType)
	fmt.Printf("%v\n", fileType)
	fmt.Printf("%v Implements %v: %t\n", fileType, writerType, fileType.Implements(writerType))
}
