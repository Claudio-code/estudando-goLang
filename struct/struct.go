package main

import (
	"fmt"
)

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (produto Produto) precoComDesconto() float64 {
	return produto.preco * (1 - produto.desconto)
}

func main() {
	produto1 := Produto{
		nome:     "carlos",
		preco:    1.43,
		desconto: 0.10,
	}

	valorDesconto := produto1.precoComDesconto()
	fmt.Println(valorDesconto)
}
