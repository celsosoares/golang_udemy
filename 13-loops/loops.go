package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	//em GO a unica estrutura de loops é o for

	//i := 0
	//tipo 1
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	//fmt.Println(i)

	//tipo 2
	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando i", j)
	// 	time.Sleep(time.Second)
	// }

	//tipo 3
	nomes := [3]string{"Celso", "Caio", "Cassio"}
	for i, nome := range nomes {
		fmt.Println(i, nome)
	}

	for i, letra := range "PALAVRA" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Celso",
		"sobrenome": "Soares",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
