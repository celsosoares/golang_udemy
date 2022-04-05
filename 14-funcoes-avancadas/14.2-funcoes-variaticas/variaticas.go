package main

import "fmt"

//recebe um numero de parametros nao definido
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	fmt.Println("Funções Variaticas")
	somaTotal := soma(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(somaTotal)
}
