package main

import (
	"errors"
	"fmt"
)

func main() {

	//NUMEROS INTEIROS

	//atribuição por inferencia
	numero0 := 0
	fmt.Println(numero0)

	//int(sem especificação) = é definida de acordo com arquitetura do computador
	//int8, int16, int32, int64 = cada tipo referente ao tamanho do inteiro recebido

	var numero int64 = -100000000000
	fmt.Println(numero)

	//nao suporta numeros negativos
	var numero2 uint = 100000000000
	fmt.Println(numero2)

	//alias = "apelido"
	//INT32 = Rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//UINT8 = Byte
	var numero4 byte = 123
	fmt.Println(numero4)

	//NUMEROS REAIS

	//atribuição por inferencia
	numero_real0 := 456.8
	fmt.Println(numero_real0)

	//Float tem q ser especificado
	var numero_real1 float32 = 123.45
	fmt.Println(numero_real1)

	//Float tem q ser especificado
	var numero_real2 float64 = 123000000000000000000000.45
	fmt.Println(numero_real2)

	//STRINGS

	//OBS: strings em GO tem q ser passado entre aspas duplas
	// aspas simples com texto nao existe em go
	// aspas simples com um caracter ele retorna o numero equivalente na tabela ASCII
	// não existe 'char' em GO

	//atribuição por inferencia
	str0 := "texto 0"
	fmt.Println(str0)

	var str string = "Texto"
	fmt.Println(str)

	str2 := 'A'
	fmt.Println(str2)

	// No GO temos valores iniciais( valor zero ) para variaveis que nao tenham valores
	// strings = string vazia " "
	// int = 0
	// float = 0.0
	// bool = false
	// error = <nil>

	var exemplo_erro error
	fmt.Println(exemplo_erro)

	var exemplo_erro2 error = errors.New("Erro interno")
	fmt.Println(exemplo_erro2)
}
