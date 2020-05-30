package constantes

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
}
