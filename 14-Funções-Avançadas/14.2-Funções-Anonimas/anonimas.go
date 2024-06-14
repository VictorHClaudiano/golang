package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá mundo")
	}()
	
	func(texto string) {
		fmt.Println(texto)
	}("passando um parametro")

}