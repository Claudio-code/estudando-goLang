package main

import "fmt"

// SLICE NÃO É UM ARRAY SLICE É PARTE DE UM ARRAY

func main() {
	a1 := [5]int{1, 2, 3, 4, 55} // array tamanho fixo

	s1 := []int{1, 2, 3} // slice tamanho variavel

	fmt.Println(a1, s1)

	// gerando slices apartir de um array

	s2 := a1[1:3]
	fmt.Println(s2)

	s3 := a1[:4]
	fmt.Println(s3)
}
