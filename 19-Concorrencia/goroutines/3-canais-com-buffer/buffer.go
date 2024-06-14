package main

import "fmt"

func main() {                   //2 é um buffer com valor 2
	canal := make(chan string,  2) //canal com buffer ele so bloqueia quando atingir a capacidade maxima dele

	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal
	
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}