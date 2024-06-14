package main

import "fmt"

// Closure sao funções que referenciam variaveis que estao fora do seu corpo

func closure() func() {
	texto := "Dentro da função Closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}