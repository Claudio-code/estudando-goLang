package main

import "fmt"

// mostrando os dois slices apontando pro msm array

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 3, 2, 1)

	fmt.Println(s1, s2)

	s1[2] = 244

	fmt.Println(s1, s2)
}
