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
}
