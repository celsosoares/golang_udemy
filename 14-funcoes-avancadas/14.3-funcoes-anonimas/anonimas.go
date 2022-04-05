package main

import "fmt"

func main() {
	fmt.Println("Funções Anonimas")

	func() {
		fmt.Println("Olá mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("texto a ser escrito")
}
