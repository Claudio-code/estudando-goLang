package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(2, 21212)
	fmt.Println("tipo =>", reflect.TypeOf(2121))

	// sem sinal (só positivos) uinit8 uinit16 uinit32 uinit64
	var bite byte = 244
	fmt.Println("tipo byte é = ", reflect.TypeOf(bite))

	// com sinal uinit8 uinit16 uinit32 uinit64
	i1 := math.MaxInt64
	fmt.Println("o valor é = ", i1)
	fmt.Println("o tipo", reflect.TypeOf(i1))

	// boolean
	boole := false
	fmt.Println("O tipo boolean é ", reflect.TypeOf(boole))
	fmt.Println(!boole)

	// string
	str := "cagada torta no mato"
	fmt.Println(str + "!")
	str2 := `
		teste
		de quebra
		de cabeça	
	`
	fmt.Println(str2, len(str2))

	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v\n", a, b, c, d, e)
}
