package main

import "fmt"

func main() {
	canal := make(chan string, 2) // 2 = quantidade total do buffer do canal, pasando disso, gera deadlock
	canal <- "Olá mundo"
	canal <- "Programando em Go"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
