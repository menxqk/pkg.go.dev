package main

import "fmt"

func main() {
	sl1 := make([]string, 10, 20)
	fmt.Printf("sl1 length: %d\n", len(sl1))
	fmt.Printf("sl1 cap: %d\n", cap(sl1))

	sl2 := make([]string, 10, 20)
	fmt.Printf("sl2 length: %d\n", len(sl2))
	fmt.Printf("sl2 cap: %d\n", cap(sl2))

	map1 := make(map[string]string, 10)
	fmt.Printf("map1 length: %d\n", len(map1))

	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("String%d", i)
		sl1[i] = s
		map1[s] = fmt.Sprintf("%d", i)
	}

	copy(sl2, sl1)

	fmt.Println(sl1)
	fmt.Println(sl2)
	fmt.Println(map1)
}
