package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) // criando o canal
	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal // o programa só continua se canal tiver algum conteudo
		// canal retorna a mensagem e o status(aberto ou fechado)
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // estrutura diferente para Canais, no caso ai, canal recebe texto
		time.Sleep(time.Second)
	}
}
