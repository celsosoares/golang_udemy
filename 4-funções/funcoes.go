package main

import "fmt"

func soma(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//quando o tipo nao é explicitado no proprio paramentro ele pega o tipo do outro parametro seguinte
//função de retorno multiplo
func calculo_matematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	result := soma(10, 20)
	fmt.Println(result)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result2 := f("Texto da função")
	fmt.Println(result2)

	// o caracter "_" (underline) é usada para caso de um retorno precisar ser ignorado
	// result_soma, _ := calculo_matematicos(10, 15)
	result_soma, result_sub := calculo_matematicos(10, 15)
	fmt.Println(result_soma, result_sub)
}
