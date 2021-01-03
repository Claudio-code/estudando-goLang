package main

import "fmt"

func main() {
	s := make([]int, 10)

	fmt.Println(s)
	s[3] = 12
	fmt.Println(s)

	s = make([]int, 20, 25)
	// len(s) pega o tamanho do slice ou do array
	// cap(s) pega a capacidade maxima do array

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 32, 12, 345, 11)
	fmt.Println(s)

	s = append(s, 12, 85)

	fmt.Println(s, len(s), cap(s))
}
