package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 4.2 // <<- inferencia do tipo em tempo de execucao

	// forma reduzida de declarar uma variavel
	area := PI * math.Pow(raio, 2)
	fmt.Println(area)

	const (
		a = 2
		b = 1
	)
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print("\n")

	var (
		c = 12
		va = 21
	)
	fmt.Print(c)
	fmt.Print(va)
	fmt.Print("\n")
	var e, f bool = true, false
	fmt.Print(e, f)
	fmt.Print("\n")

	g, h, i := 2, "cagada", false
	fmt.Print(g, h, i)
}
