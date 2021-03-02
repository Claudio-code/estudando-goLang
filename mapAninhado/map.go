package main

import "fmt"

func main() {
	casas := make(map[int]string)
	ruas := make(map[string]map[int]string)

	casas[12] = "joana"
	casas[32] = "marcos"
	casas[534] = "alesanddro"

	ruas["rua iguacu"] = casas
	ruas["b"] = casas
	fmt.Println(ruas)

	delete(ruas, "b")

	fmt.Println(ruas)
}
