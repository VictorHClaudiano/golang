package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5] string
	array1[0] = "posição 1"
	fmt.Println(array1)

	array2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)


	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posiceção Alterada"
	fmt.Println(slice2)


	//Arrays Internos
	//função make aloca um espaço na memoria
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	
	slice3 = append(slice3, 5)  //slice nunca tem um tamanho limite e quando passa ele dobra a quantidade
	slice3 = append(slice3, 6)
	
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4)) 
}

//Slice é um estrutura feita baseada no array com flexibilidade a mais.
//Slice nao é um array, apenas parece mais não é.
//append adiciona ou altera valor a um slice.
//slice é um pedaço de um array