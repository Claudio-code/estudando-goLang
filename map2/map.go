package main

import "fmt"

func main() {
	funcionariosSalarios := map[string]float64{
		"Ricardo": 123,
		"Diego":   345,
		"Diogo":   431,
		"cidemar": 876,
	}

	funcionariosSalarios["carla bassi"] = 12312
	fmt.Println(funcionariosSalarios)
}
