package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte{0, 10, 20, 30, 40}
	b := []byte{30, 40, 50, 60, 10}
	c := make([]byte, len(a), cap(a))
	fmt.Printf("copied: %d\n", copy(c, a))
	d := bytes.Join([][]byte{a, c}, []byte{99, 99})

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)

	fmt.Printf("Compare a,b: %v\n", bytes.Compare(a, b))
	fmt.Printf("Compare a,c: %v\n", bytes.Compare(a, c))

	fmt.Printf("a Contains rune '\\n' (%d): %t\n", '\n', bytes.ContainsRune(a, '\n'))
	fmt.Printf("Position of rune'\\n' in a: %d\n", bytes.Index(a, []byte{'\n'}))
	fmt.Printf("b Contains rune '\\n' (%d): %t\n", '\n', bytes.ContainsRune(b, '\n'))
	fmt.Printf("Position of rune'\\n' in b: %d\n", bytes.Index(b, []byte{'\n'}))

	const nihongo = "日本語"
	fmt.Printf("Convert %s to bytes: %v\n", nihongo, []byte(nihongo))
	fmt.Printf("Convert %s to runes: %v\n", nihongo, bytes.Runes([]byte(nihongo)))
	fmt.Println()

	buf := bytes.NewBuffer([]byte(nihongo))
	fmt.Printf("buf length: %d, buf size: %d\n", buf.Len(), buf.Len())
	buf.WriteString(" ABC")
	fmt.Printf("%s\n", buf)

	reader := bytes.NewReader(d)
	fmt.Printf("reader length: %d, reader size: %d\n", reader.Len(), reader.Size())

}
