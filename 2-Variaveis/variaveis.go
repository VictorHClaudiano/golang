package main

import "fmt"
//2 exemplos de chamar uma função 
//uma usando o var e chamando o tipo de variavel que vai ser usado!
//a segunda seria inferencia de tipo (baseado no valor dele)
func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "variavel 2"
	
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5, variavel6)
}
