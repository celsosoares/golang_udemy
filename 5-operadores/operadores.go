package main

import "fmt"

func main() {
	//ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 1 / 2
	resto := 1 % 2
	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	var numero1 int16 = 10
	var numero2 int32 = 25
	//qualquer operação com tipos diferentes nao da certo (msm sendo numeros ou strings)
	// exemplo: int16 com int16, int8 com int8, float21 com float32, ...
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 <= 2)
	//...
	//...
	//retorna true ou false

	//OPERADORES LOGICOS
	fmt.Println(true && true)  //and
	fmt.Println(true || false) //or
	fmt.Println(!true)         //not

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	numero--
	//...
	fmt.Println(numero)
}
