package main

import (
	"fmt"
	"mime"
)

func main() {
	mediaType := "text/html"
	params := map[string]string{
		"charset": "utf-8",
	}

	result := mime.FormatMediaType(mediaType, params)

	fmt.Println("result:", result)
}
