package main

import "fmt"

func main() {
	//Aritimeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)


	//ATRIBUIÇÃO

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//Operadores RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 < 2)


	//Operadores LOGICOS
	fmt.Println("-----------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro &&  falso) // And
	fmt.Println(verdadeiro || falso) // OU ou OR
	fmt.Println(!verdadeiro) //Negação
	
	//Operadores UNARIOS
	numero := 10
	numero ++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero --
	numero -= 20 //numero = numero -20

	numero *= 3 //numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
	
	//Operador TERNÁRIO

	var texto string
	if numero > 5 {
		texto = "maior que 5"
	}else {
		texto = "menor que 5"
	}
	fmt.Println(texto)
}

