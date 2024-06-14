package main
//Podemos ter uma variatica por função e tem que estar localizado no ultima coisa a ser declarada
import "fmt"

func soma(numeros ...int) int{ //...int é uma variatica
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalDaSoma := soma (1, 2, 3, 4, 5, 6, 7, 100, 2020, 30)
	fmt.Println(totalDaSoma)
}