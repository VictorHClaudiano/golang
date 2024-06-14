package main

import "fmt"
//Init é função que vai ser executada antes da função main
//Só pode ter uma(1) por aquivo. OBS : Não uma por pacote

func init() {
	fmt.Println("Executando função init")
}


func main() {
	fmt.Println("Função main sendo executada")
}