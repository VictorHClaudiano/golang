package main

import (
	"fmt"
	//"time"
)

func main() {
//	i := 0

//	for i < 10 {
//		i++
//		fmt.Println("Incrementano i")
//		time.Sleep(time.Second)
	//}

	//fmt.Println(i)

	//for j := 0; j < 10; j++ {
	//	fmt.Println("Incrementano j", j)
	//	time.Sleep(time.Second)
	//}

	nomes := [3]string{"Jão", "Davi", "Lucas"}
	
	for indice, nome := range nomes{
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string {
		"nome":		"Victor",
		"sobrenome":"Hugo",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}


//*********************NÃO DA PARA FAZER RANGE EM STRUCTS*****************