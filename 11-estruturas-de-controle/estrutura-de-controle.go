package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Dados")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero == 9 {
		fmt.Println("O numero é 9")
	} else {
		fmt.Println("O numero é menor que 15 e nao é 9")
	}

	//if init (criando uma varivel dentro do if)
	// só é possivel utilizar essa variavel dentro do escopo do if-else
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	}

}
