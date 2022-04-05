package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

//metodo

/*
Um método é uma função com um receptor. Uma declaração de método vincula um identificador e o nome do método a um método e associa o método com o tipo base do
receptor. Métodos são muito parecidos com funções, mas são chamados invocando-os em uma instância de um tipo específico.
Enquanto você chama funções onde quiser, como por exemplo em Area(retangulo), você só pode chamar métodos em "coisas" específicas.

A sintaxe para definir um método é similar à sintaxe para definir uma função. A única diferença é a adição de um parâmetro extra após a palavra-chave func,
para especificar o receptor do método. O receptor é uma declaração do tipo que você deseja para definir o método.
*/

// obrigatoriamente esta associado a algo, alguma estrutura
// func ("variavel" "estrutura ao qual ta associado") ("nome do metodo()")
func (u usuario) salvar() { //entre parenteses se passa a estrutura que esta referenciando
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maior_de_idade() bool {
	return u.idade >= 18
}

func main() {
	usuario1 := usuario{"Celso", 24}
	usuario1.salvar()

	usuario2 := usuario{"Caio", 32}
	usuario2.salvar()

	maiorDeIdade := usuario2.maior_de_idade()
	fmt.Println(maiorDeIdade)

}
