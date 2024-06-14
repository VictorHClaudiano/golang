package main
//utilizamos o generation para esconder a complexidade
//encapsula uma chamada de uma go routine e retorna um canal
import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal<- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	
	return canal
}