package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func main() {
	// DEFER = ADIAR
	// ele só executa a função que tiver o defer antes do retorno
	// muito usado em banco de dados
	fmt.Println("Defer")

	defer funcao1()
	funcao2()
}
