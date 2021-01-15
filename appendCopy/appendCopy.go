package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	slice1 = append(slice1, 4, 5, 23)
	fmt.Println(slice1, array1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)

	fmt.Println(slice2)
}