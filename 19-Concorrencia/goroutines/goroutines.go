package main

import (
	"fmt"
	"time"
)

//CONCORRENCIA != PARALELISMO(duas ou mais tarefas que estao sendo executada ao mesmo tempo)
//CONCORENCIA (Elas podem ou nao estar sendo executada ao mesmo tempo)

func main() {
	go escrever("Olá Mundo!") //goroutine sao extramamentes eficientes e voce pode ter milhares dentro do codigo
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}