package main

//canal pois pode tanto enviar como receber dados
import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) //chan palavra chave e o seu tipo que so vai ser aceito no tipo especificado
	go escrever("Olá Mundo!", canal)

	mensagem := <-canal // com a seta primeiro voce esta esperando receber um valor
	fmt.Println(mensagem)
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto  // a seta despois do canal esta mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}
}