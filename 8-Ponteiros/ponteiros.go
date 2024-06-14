package main

import "fmt"
// PONTEIRO É UMA REFERENCIA DE MEMORIA 
//salvar dentro dele o endereço de memoria
func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	
	fmt.Println(variavel1, variavel2)
	
	variavel1 ++
	fmt.Println(variavel1, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) //usar * (desreferenciação) para saber o valor que esta no ponteiro (endereço de memoria)
}