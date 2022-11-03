package main

import (
	"fmt"
	"mime"
)

func main() {
	mediaType, params, err := mime.ParseMediaType("text/html; charset=utf-8")
	if err != nil {
		panic(err)
	}

	fmt.Println("type:", mediaType)
	fmt.Println("charset:", params["charset"])
}
