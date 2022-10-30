package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	goodJSON := `{"example": 1}`
	badJSON := `{"exaple":2:]}}`

	fmt.Println(json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))
}
