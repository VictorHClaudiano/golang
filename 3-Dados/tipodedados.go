package main

import (
	"errors"
	"fmt"
)

//O que muda são a taxa de bits int8, int16, int32, int64
func main() {
	numero := 10000000000000
	fmt.Println(numero)

	//uint não aceita numeros negativos
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	// int32 = rune
	var numero3 rune = 12456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)


	//float aceita numeros quebrados
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)
	
	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// Fim dos numeros reais

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := '%' //converte em numero da tabela ASK
	fmt.Println(char)
	// Fim Strings

	//valor inicial string (vazio) int e float (zero) boolean(false) erro (nil)
	var texto string
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error =  errors.New("Erro interno")
	fmt.Println(erro)
}
