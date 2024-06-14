package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controle")

	numero := -5
//Você cria para usar dentro de um escopo apenas.
	if numero > 15 {
		fmt.Println("maior que 15")
	} else{
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que zero")
	} else if numero < - 10{
		fmt.Println("Numero é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}




}