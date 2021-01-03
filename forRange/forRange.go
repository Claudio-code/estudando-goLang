package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5, 21}

	for i, _ := range numeros {
		fmt.Println(i)
	}

	for _, item := range numeros {
		fmt.Println(item)
	}
}
