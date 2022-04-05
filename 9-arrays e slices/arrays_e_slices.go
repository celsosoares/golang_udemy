package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"string1", "string2", "string3", "string4", "string5"}
	fmt.Println(array2)

	//[...] array indefinido, ele gera o tamanho a partir dos dados de entrada
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//SLICE (é mais flexivel que o array) - tamanho dinamico
	slice := []int{10, 11, 12, 13}
	fmt.Println(slice)

	slice = append(slice, 14)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posicao alterada"
	fmt.Println(slice2)

	// Arrays Internos

	//:make: aloca um espaço na memoria pra determinado dado do programa
	slice3 := make([]float32, 10, 15) //tipo, tamanho, tamanho max(se omitir esse parametro ele considera o tamanho normal como tamanho max)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	// se o slice estourar sua capacidade maxima, ele dobra o tamanho do slice

}
