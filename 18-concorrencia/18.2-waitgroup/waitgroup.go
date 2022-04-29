package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //quantidade de goroutines a ser executadas

	go func() { //função anonima
		escrever("Olá mundo!")
		waitGroup.Done() // -1 gorotines
	}()
	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1 gorotines
	}()

	waitGroup.Wait() // ate os gorotines chegar em 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
