package main

import "fmt"

var n int

// vai ser executada antes da função main
// independente da ordem
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}
