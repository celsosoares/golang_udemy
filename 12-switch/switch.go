package main

import "fmt"

func dia_da_semana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Numero Inválido"
	}
}

//outra forma de usar o switch
func dia_da_semana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Numero Inválido"
	}
}

//outra forma de usar o switch
func dia_da_semana3(numero int) string {
	var dia_da_semana string

	switch {
	case numero == 1:
		dia_da_semana = "Domingo"
	case numero == 2:
		dia_da_semana = "Segunda-Feira"
	case numero == 3:
		dia_da_semana = "Terça-Feira"
	case numero == 4:
		dia_da_semana = "Quarta-Feira"
	case numero == 5:
		dia_da_semana = "Quinta-Feira"
	case numero == 6:
		dia_da_semana = "Sexta-Feira"
	case numero == 7:
		dia_da_semana = "Sabado"
	default:
		dia_da_semana = "Numero Inválido"
	}
	return dia_da_semana
}

func main() {
	fmt.Println("Switch")
	fmt.Println(dia_da_semana(9))
}
