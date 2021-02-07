package main

import "fmt"

func main() {
	// maps tem q ser inicializados
	aprovados := make(map[int]string)

	aprovados[123] = "maria"
	aprovados[321] = "claudia"
	aprovados[456] = "marcos"
	fmt.Println(aprovados)

	fmt.Println(aprovados[123] + "\n\n")

	aprovados[123] = "jeoconda"
	for cpf, nome := range aprovados {
		fmt.Println(cpf)
		fmt.Println(nome)
	}

	delete(aprovados, 123)
	fmt.Println(aprovados[123] + "\n\n")
}
