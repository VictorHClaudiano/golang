package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup //pacote do go sync 

	waitGroup.Add(2) //2 goroutines que estao em espera

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() //metodo done tira um do contador -1
	}()

	 go func(){
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	 }()
	 
	 waitGroup.Wait() //fala para função main esperar a contagem da goroutines chegar em 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}