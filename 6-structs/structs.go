package main

import "fmt"

/*
Structs são uma coleção de campos digitados. Basicamente, um grupo de informações é agrupado como um tipo,
que pode ser usado para criar instâncias da estrutura que definimos.
*/
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco //exemplo com struct dentro de struct
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivos structs")

	enderecoExemplo := endereco{"Rua 1", 122}

	var u usuario
	u.nome = "celso"
	u.idade = 21
	fmt.Println(u)

	u2 := usuario{"davi", 24, enderecoExemplo}
	fmt.Println(u2)

	//qnd nao se tem todos os campos necessarios, se especifica o campo
	u3 := usuario{nome: "davi"}
	fmt.Println(u3)

}
