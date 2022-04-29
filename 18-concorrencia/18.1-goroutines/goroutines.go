package main

import (
	"fmt"
	"time"
)

// "go" palavra chave usada pra "fugir" da rotina, ou seja, o comando nao espera ser finalizado para ir para a proxima linha

func main() {
	go escrever("Olá mundo!") //goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
